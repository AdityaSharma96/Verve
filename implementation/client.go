package implementation

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type PostData struct {
	UniqueRequestCount int `json:"unique_request_count"`
}

func sendPostRequest(endpoint string, count int) {
	data := PostData{UniqueRequestCount: count}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("HTTP POST failed: %v", err)
		return
	}
	log.Printf("HTTP POST to %s returned status %d", endpoint, resp.StatusCode)
	resp.Body.Close()
}
