package implementation

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

var uniqueRequests sync.Map

func LogRequestCount(kafkaWriter *kafka.Writer) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		<-ticker.C
		count := getCurrentMinuteCount()
		log.Printf("Unique requests in the last minute: %d", count)

		message := kafka.Message{
			Key:   []byte("unique-count"),
			Value: []byte(strconv.Itoa(count)),
		}
		if err := kafkaWriter.WriteMessages(context.Background(), message); err != nil {
			log.Printf("Failed to write to Kafka: %v", err)
		}

		uniqueRequests = sync.Map{}
	}
}

// getCurrentMinuteCount returns the number of unique requests in the current minute
func getCurrentMinuteCount() int {
	count := 0
	uniqueRequests.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	return count
}
