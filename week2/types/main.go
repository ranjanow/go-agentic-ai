package main

import "fmt"

// --- Boolean types ---
func booleanExamples() {
	var a bool = true
	b := false
	fmt.Println("Boolean:", a, b, a && b, a || b, !a)
}

// --- Numeric types ---
func numericExamples() {
	var i8 int8 = 127
	var u32 uint32 = 4294967295
	var f64 float64 = 3.14159265358979
	var c128 complex128 = complex(2, 3)

	fmt.Printf("int8: %d, uint32: %d\n", i8, u32)
	fmt.Printf("float64: %.6f\n", f64)
	fmt.Printf("complex128: %v (real=%.0f, imag=%.0f)\n", c128, real(c128), imag(c128))
}

// --- String type ---
func stringExamples() {
	s := "Hello, Go spec!"
	raw := `raw string: no \n escape here`

	fmt.Println(s)
	fmt.Println(raw)
	fmt.Println("Length in bytes:", len(s))
	fmt.Println("First byte:", s[0]) // byte value
}

// --- Array types (length is part of the type) ---
func arrayExamples() {
	var arr [5]int
	arr[2] = 42
	fmt.Println("Array:", arr)

	// [3]int and [5]int are DIFFERENT types
	a3 := [3]int{1, 2, 3}
	fmt.Println("3-element array:", a3)
}

// --- Slice types ---
func sliceExamples() {
	// Slice: dynamic view into an array
	s := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice:", s, "len:", len(s), "cap:", cap(s))

	// Sub-slice
	fmt.Println("s[1:3]:", s[1:3])

	// Append grows the slice
	s = append(s, 60, 70)
	fmt.Println("After append:", s)

	// make: create a slice with length and capacity
	made := make([]string, 3, 5)
	made[0] = "agent"
	made[1] = "tool"
	made[2] = "memory"
	fmt.Println("made:", made)
}

// --- Struct types ---
type Agent struct {
	Name   string
	Model  string
	Tools  []string
}

func structExamples() {
	a := Agent{
		Name:  "GoAgent",
		Model: "claude-sonnet",
		Tools: []string{"web_search", "code_exec"},
	}
	fmt.Printf("Agent: %s using %s with tools %v\n", a.Name, a.Model, a.Tools)
}

// --- Pointer types ---
func pointerExamples() {
	n := 99
	ptr := &n
	fmt.Printf("Pointer: address=%p value=%d\n", ptr, *ptr)
	*ptr = 100
	fmt.Println("After modification via pointer:", n)
}

// --- Interface types ---
type Speaker interface {
	Speak() string
}

type GoAgent struct{ Name string }
type PythonAgent struct{ Name string }

func (g GoAgent) Speak() string     { return g.Name + " speaks in Go" }
func (p PythonAgent) Speak() string { return p.Name + " speaks in Python" }

func interfaceExamples() {
	agents := []Speaker{
		GoAgent{Name: "Gopher"},
		PythonAgent{Name: "Snek"},
	}
	for _, a := range agents {
		fmt.Println(a.Speak())
	}
}

// --- Map types ---
func mapExamples() {
	// Maps are reference types; zero value is nil
	m := map[string]int{
		"goroutines": 10,
		"channels":   5,
		"interfaces": 3,
	}
	fmt.Println("Map:", m)

	// Comma-ok idiom for safe lookup
	if v, ok := m["goroutines"]; ok {
		fmt.Println("goroutines value:", v)
	}

	// Delete a key
	delete(m, "interfaces")
	fmt.Println("After delete:", m)
}

// --- Channel types ---
func channelExamples() {
	ch := make(chan string, 2) // buffered channel

	ch <- "message from Go spec study"
	ch <- "channels enable safe goroutine communication"

	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
}

func main() {
	fmt.Println("=== Week 2: Go Type System (from go.dev/ref/spec) ===\n")

	fmt.Println("-- Boolean types --")
	booleanExamples()

	fmt.Println("\n-- Numeric types --")
	numericExamples()

	fmt.Println("\n-- String types --")
	stringExamples()

	fmt.Println("\n-- Array types --")
	arrayExamples()

	fmt.Println("\n-- Slice types --")
	sliceExamples()

	fmt.Println("\n-- Struct types --")
	structExamples()

	fmt.Println("\n-- Pointer types --")
	pointerExamples()

	fmt.Println("\n-- Interface types --")
	interfaceExamples()

	fmt.Println("\n-- Map types --")
	mapExamples()

	fmt.Println("\n-- Channel types --")
	channelExamples()
}
