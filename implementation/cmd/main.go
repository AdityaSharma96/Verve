package main

import (
	"Verve/implementation"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
)

var (
	rdb         *redis.Client
	kafkaWriter *kafka.Writer
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "unique-request-counts",
		Balancer: &kafka.LeastBytes{},
	}
}

func main() {
	go implementation.LogRequestCount(kafkaWriter)

	http.HandleFunc("/api/verve/accept", func(w http.ResponseWriter, r *http.Request) {
		implementation.AcceptHandler(rdb, w, r)
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
