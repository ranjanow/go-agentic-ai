package main

import (
	"encoding/json"
	"fmt"
)

// Message represents an Anthropic-style chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Request represents an API request body
type Request struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []Message `json:"messages"`
}

// Choice is part of a response
type Choice struct {
	Text         string `json:"text"`
	FinishReason string `json:"finish_reason"`
}

// APIResponse is a simplified API response
type APIResponse struct {
	ID      string   `json:"id"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

func main() {
	fmt.Println("=== Week 3: encoding/json ===\n")

	// --- Marshal: Go struct → JSON ---
	req := Request{
		Model:     "claude-sonnet-4-6",
		MaxTokens: 1024,
		Messages: []Message{
			{Role: "user", Content: "Explain goroutines in Go."},
		},
	}

	jsonBytes, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println("-- Marshalled JSON --")
	fmt.Println(string(jsonBytes))

	// --- Unmarshal: JSON → Go struct ---
	jsonStr := `{
		"id": "msg_01abc",
		"model": "claude-sonnet-4-6",
		"choices": [
			{"text": "A goroutine is a lightweight thread managed by the Go runtime.", "finish_reason": "stop"}
		]
	}`

	var apiResp APIResponse
	if err := json.Unmarshal([]byte(jsonStr), &apiResp); err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Println("\n-- Unmarshalled struct --")
	fmt.Printf("ID:     %s\n", apiResp.ID)
	fmt.Printf("Model:  %s\n", apiResp.Model)
	fmt.Printf("Answer: %s\n", apiResp.Choices[0].Text)

	// --- map[string]any for flexible JSON ---
	flexible := `{"name": "GoAgent", "version": 1, "active": true}`
	var m map[string]any
	json.Unmarshal([]byte(flexible), &m)
	fmt.Println("\n-- Flexible map decode --")
	for k, v := range m {
		fmt.Printf("  %s: %v (%T)\n", k, v, v)
	}
}
