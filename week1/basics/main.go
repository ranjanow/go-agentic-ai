package main

import "fmt"

// Intern represents a person in the internship program
type Intern struct {
	Name         string
	Organization string
	Week         int
}

// Greet returns a greeting string for the intern
func (i Intern) Greet() string {
	return fmt.Sprintf("Hi, I'm %s at %s — Week %d", i.Name, i.Organization, i.Week)
}

// add returns the sum of two integers
func add(a, b int) int {
	return a + b
}

// fibonacci returns the first n Fibonacci numbers
func fibonacci(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		if i < 2 {
			nums[i] = i
		} else {
			nums[i] = nums[i-1] + nums[i-2]
		}
	}
	return nums
}

func main() {
	// Variables and short declaration
	language := "Go"
	version := 1.22
	fmt.Printf("Learning %s %.2f\n", language, version)

	// Struct usage
	intern := Intern{Name: "Your Name", Organization: "AgenticGokit", Week: 1}
	fmt.Println(intern.Greet())

	// Function call
	fmt.Println("3 + 5 =", add(3, 5))

	// Loop + slice
	fmt.Println("First 8 Fibonacci numbers:", fibonacci(8))

	// Map
	topics := map[string]string{
		"week1": "Tour of Go",
		"week2": "Go Language Spec",
		"week3": "HTTP & APIs",
		"week4": "Agentic AI",
	}
	for week, topic := range topics {
		fmt.Printf("  %s → %s\n", week, topic)
	}
}
