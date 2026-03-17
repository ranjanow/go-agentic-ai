package main

import (
	"context"
	"fmt"
	"time"
)

// Message represents a chat message
type Message struct {
	Role    string
	Content string
}

// Tool represents a callable tool the agent can use
type Tool struct {
	Name        string
	Description string
	Execute     func(input string) (string, error)
}

// Agent is the core agentic loop struct
type Agent struct {
	Name     string
	Model    string
	Tools    map[string]Tool
	History  []Message
	MaxSteps int
}

// NewAgent creates a new agent with default settings
func NewAgent(name, model string) *Agent {
	return &Agent{
		Name:     name,
		Model:    model,
		Tools:    make(map[string]Tool),
		History:  []Message{},
		MaxSteps: 10,
	}
}

// RegisterTool adds a tool to the agent's toolkit
func (a *Agent) RegisterTool(t Tool) {
	a.Tools[t.Name] = t
	fmt.Printf("[agent] registered tool: %s\n", t.Name)
}

// AddMessage appends a message to the conversation history
func (a *Agent) AddMessage(role, content string) {
	a.History = append(a.History, Message{Role: role, Content: content})
}

// Think simulates an LLM call (placeholder — replace with real API call in Week 4)
func (a *Agent) Think(ctx context.Context, userInput string) (string, error) {
	// In a real implementation, this would call the Anthropic Messages API:
	//
	//   POST https://api.anthropic.com/v1/messages
	//   {
	//     "model": a.Model,
	//     "max_tokens": 1024,
	//     "messages": a.History
	//   }
	//
	// For now, we simulate a response
	select {
	case <-ctx.Done():
		return "", fmt.Errorf("think cancelled: %w", ctx.Err())
	case <-time.After(200 * time.Millisecond): // simulate API latency
		return fmt.Sprintf("[simulated LLM response to: %q]", userInput), nil
	}
}

// Run executes the agent loop for a given task
func (a *Agent) Run(ctx context.Context, task string) error {
	fmt.Printf("\n[agent:%s] starting task: %s\n", a.Name, task)
	a.AddMessage("user", task)

	for step := 1; step <= a.MaxSteps; step++ {
		fmt.Printf("[agent] step %d/%d\n", step, a.MaxSteps)

		// 1. Think: call LLM with current history
		response, err := a.Think(ctx, task)
		if err != nil {
			return fmt.Errorf("step %d think failed: %w", step, err)
		}
		a.AddMessage("assistant", response)
		fmt.Printf("[agent] response: %s\n", response)

		// 2. Check if task is done (simple heuristic — real agents parse tool calls)
		// In a full implementation: parse response for tool_use blocks,
		// execute tools, append results, and loop until stop_reason == "end_turn"
		fmt.Printf("[agent] task complete after %d step(s)\n", step)
		break
	}

	return nil
}

func main() {
	fmt.Println("=== Week 4: Agent Loop (AgenticGokit) ===\n")

	// Create agent
	agent := NewAgent("GoAgent-v1", "claude-sonnet-4-6")

	// Register example tools
	agent.RegisterTool(Tool{
		Name:        "web_search",
		Description: "Search the web for current information",
		Execute: func(input string) (string, error) {
			return fmt.Sprintf("[search results for: %s]", input), nil
		},
	})

	agent.RegisterTool(Tool{
		Name:        "code_runner",
		Description: "Execute Go code snippets",
		Execute: func(input string) (string, error) {
			return fmt.Sprintf("[executed: %s]", input), nil
		},
	})

	fmt.Printf("\nAgent: %s\nModel: %s\nTools: %d registered\n",
		agent.Name, agent.Model, len(agent.Tools))

	// Run the agent
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := agent.Run(ctx, "Explain how goroutines differ from OS threads")
	if err != nil {
		fmt.Println("Agent error:", err)
		return
	}

	// Print conversation history
	fmt.Println("\n-- Conversation history --")
	for _, msg := range agent.History {
		fmt.Printf("[%s]: %s\n", msg.Role, msg.Content)
	}
}
