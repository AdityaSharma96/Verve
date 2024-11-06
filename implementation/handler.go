package implementation

import (
	"context"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

func AcceptHandler(rdb *redis.Client, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	endpoint := r.URL.Query().Get("endpoint")

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	if err := recordUniqueRequest(rdb, id); err != nil {
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}

	if endpoint != "" {
		go func() {
			count := getCurrentMinuteCount()
			sendPostRequest(endpoint, count)
		}()
	}

	w.Write([]byte("ok"))
}

func recordUniqueRequest(rdb *redis.Client, id string) error {
	ctx := context.Background()
	// Set the key with a 1-minute TTL if it doesn't exist already
	return rdb.SetNX(ctx, id, true, time.Minute).Err()
}
