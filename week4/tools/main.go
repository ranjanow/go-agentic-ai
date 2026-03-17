package main

import (
	"errors"
	"fmt"
	"strings"
)

// ToolInput is the input passed to a tool
type ToolInput struct {
	Name  string
	Param string
}

// ToolOutput is the result from a tool
type ToolOutput struct {
	ToolName string
	Result   string
	Error    error
}

// ToolFunc is the signature for all tool functions
type ToolFunc func(param string) (string, error)

// Registry holds all available tools
type Registry struct {
	tools map[string]ToolFunc
}

// NewRegistry creates an empty tool registry
func NewRegistry() *Registry {
	return &Registry{tools: make(map[string]ToolFunc)}
}

// Register adds a tool to the registry
func (r *Registry) Register(name string, fn ToolFunc) {
	r.tools[name] = fn
	fmt.Printf("[registry] registered: %s\n", name)
}

// Dispatch executes a tool by name
func (r *Registry) Dispatch(input ToolInput) ToolOutput {
	fn, ok := r.tools[input.Name]
	if !ok {
		return ToolOutput{
			ToolName: input.Name,
			Error:    fmt.Errorf("tool %q not found in registry", input.Name),
		}
	}
	result, err := fn(input.Param)
	return ToolOutput{ToolName: input.Name, Result: result, Error: err}
}

// List returns all registered tool names
func (r *Registry) List() []string {
	names := make([]string, 0, len(r.tools))
	for name := range r.tools {
		names = append(names, name)
	}
	return names
}

// --- Example tool implementations ---

func webSearch(query string) (string, error) {
	if query == "" {
		return "", errors.New("query cannot be empty")
	}
	return fmt.Sprintf("Search results for %q: [result1, result2, result3]", query), nil
}

func calculator(expr string) (string, error) {
	// Very simple: only handles "N + M"
	parts := strings.Split(expr, "+")
	if len(parts) != 2 {
		return "", fmt.Errorf("calculator only supports addition, got: %s", expr)
	}
	var a, b float64
	fmt.Sscanf(strings.TrimSpace(parts[0]), "%f", &a)
	fmt.Sscanf(strings.TrimSpace(parts[1]), "%f", &b)
	return fmt.Sprintf("%.2f", a+b), nil
}

func summarise(text string) (string, error) {
	words := strings.Fields(text)
	if len(words) == 0 {
		return "", errors.New("no text to summarise")
	}
	preview := words
	if len(words) > 5 {
		preview = words[:5]
	}
	return fmt.Sprintf("Summary: %s... (%d words total)", strings.Join(preview, " "), len(words)), nil
}

func main() {
	fmt.Println("=== Week 4: Tool Registry & Dispatch ===\n")

	// Build registry
	reg := NewRegistry()
	reg.Register("web_search", webSearch)
	reg.Register("calculator", calculator)
	reg.Register("summarise", summarise)

	fmt.Println("\nRegistered tools:", reg.List())

	// Simulate tool calls from an agent
	calls := []ToolInput{
		{Name: "web_search", Param: "Go agentic AI frameworks 2025"},
		{Name: "calculator", Param: "42 + 58"},
		{Name: "summarise", Param: "Go is a statically typed compiled language designed at Google"},
		{Name: "unknown_tool", Param: "will fail"},
	}

	fmt.Println("\n-- Dispatching tool calls --")
	for _, call := range calls {
		out := reg.Dispatch(call)
		if out.Error != nil {
			fmt.Printf("[%s] ERROR: %v\n", out.ToolName, out.Error)
		} else {
			fmt.Printf("[%s] %s\n", out.ToolName, out.Result)
		}
	}
}
