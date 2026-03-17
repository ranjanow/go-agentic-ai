package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Response represents a generic JSON API response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// httpGet performs a GET request and returns the body bytes
func httpGet(url string, timeoutSec int) ([]byte, int, error) {
	client := &http.Client{
		Timeout: time.Duration(timeoutSec) * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("GET %s failed: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("reading body failed: %w", err)
	}

	return body, resp.StatusCode, nil
}

func main() {
	fmt.Println("=== Week 3: net/http client ===\n")

	// Example: call a public API
	url := "https://httpbin.org/get"
	fmt.Printf("GET %s\n", url)

	body, status, err := httpGet(url, 10)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Status: %d\n", status)
	fmt.Printf("Body (first 200 chars): %s\n", string(body[:min(200, len(body))]))

	// Demonstrate JSON decode
	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("JSON decode error:", err)
		return
	}
	fmt.Println("\nParsed origin:", result["origin"])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
