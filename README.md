# Verve Service

This is a Go-based REST service that processes unique requests, stores them in Redis, and sends unique request counts to a Kafka topic. The service is designed to handle at least 10,000 requests per second and includes the following features:

- **Unique request tracking** using Redis.
- **Periodic logging** of unique request counts every minute.
- **Support for an optional external endpoint** to trigger HTTP POST requests with the current unique request count.
- **Kafka integration** to send request counts to a Kafka topic.
- **Support for multiple instances** (in case of load balancing).

## Features

- **GET /api/verve/accept**:
    - Accepts a mandatory `id` query parameter (integer).
    - Optionally accepts an `endpoint` query parameter (string) for triggering external HTTP requests.
    - Returns `ok` if successful or `failed` if any error occurs.

- **Periodic logging** of unique requests every minute.
- **Kafka integration** to send the count of unique requests to a Kafka topic.
- **Redis integration** for deduplication and storage of unique request IDs.

## Prerequisites

- **Go** (version 1.19 or later): Install from [Go's official website](https://golang.org/dl/).
- **Redis**: Install Redis or use Docker to run Redis locally.
- **Kafka**: Install Kafka and Zookeeper or use Docker to run them locally.
- **Docker** (optional): To containerize the application.

## Running the Service Locally

### Step 1: Install Redis and Kafka

1. **Start Redis** using Docker (if not already installed):
   ```bash
   docker run --name redis -p 6379:6379 -d redis
   
2. Start Zookeeper (required by Kafka):
   ```bash
   docker run -d --name=zookeeper -p 2181:2181 wurstmeister/zookeeper

3. Start Kafka:
   ```bash
   docker run -d --name=kafka -p 9092:9092 --link zookeeper -e KAFKA_ADVERTISED_LISTENER=PLAINTEXT://localhost:9092 -e KAFKA_LISTENER=PLAINTEXT://localhost:9092 wurstmeister/kafka

### Step 2: Build the project
    go build -o verve-service

### Step 3: Run the service
    go run main.go

### Step 4: Test
1. Test without an endpoint:
    ```bash
   curl "http://localhost:8080/api/verve/accept?id=123"

2. Test with an endpoint:
   ```bash
   curl "http://localhost:8080/api/verve/accept?id=123&endpoint=http://example.com"
