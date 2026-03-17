package main

import "fmt"

// increment modifies the value at the pointer address
func increment(n *int) {
	*n++
}

// newCounter returns a pointer to a new int counter
func newCounter(start int) *int {
	c := start
	return &c
}

func main() {
	// Basic pointer usage
	x := 10
	p := &x
	fmt.Println("Value:", x)
	fmt.Println("Pointer address:", p)
	fmt.Println("Value via pointer:", *p)

	// Modifying via pointer
	increment(&x)
	fmt.Println("After increment:", x)

	// Pointer returned from function
	counter := newCounter(0)
	for i := 0; i < 5; i++ {
		increment(counter)
	}
	fmt.Println("Counter after 5 increments:", *counter)

	// Nil pointer check
	var q *int
	if q == nil {
		fmt.Println("q is nil — always check before dereferencing!")
	}
}
