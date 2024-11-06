# Thought Process

## Architecture and Design

1. **Concurrency**:
    - Use Go's `net/http` package for high-performance request handling.
    - Redis with `SetNX` ensures unique request IDs across multiple instances.
    - Kafka is used for reliable, distributed logging of unique request counts.

2. **Load Balancing and Deduplication**:
    - Redis is employed for deduplication across instances, ensuring unique IDs even under load balancing.

3. **Fault Tolerance**:
    - Error handling and asynchronous requests ensure the application is fault-tolerant.
    - Logging failures to Kafka or POST endpoints do not disrupt primary service flow.

4. **Extensions**:
    - **POST Request**: Uses JSON with a unique count field.
    - **Multi-instance Deduplication**: Redis provides cross-instance uniqueness.
    - **Distributed Logging**: Kafka serves as a scalable, distributed log for unique request counts.

5. **Containerization**:
    - Dockerfile provided to allow for scalable and deployable service instances.

## Assumptions
- Redis is running on `localhost:6379` for deduplication.
- Kafka is available on `localhost:9092` for distributed logging.
- Service port is set to 8080.

