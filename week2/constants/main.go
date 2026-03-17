package main

import "fmt"

// Typed constant
const MaxTokens int = 4096

// Untyped constant — takes type from context
const Pi = 3.14159265358979

// iota: auto-incrementing constant generator
type Weekday int

const (
	Sunday Weekday = iota // 0
	Monday                // 1
	Tuesday               // 2
	Wednesday             // 3
	Thursday              // 4
	Friday                // 5
	Saturday              // 6
)

func (d Weekday) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}

// iota with bit shifts — useful for flag/permission types
type Permission uint

const (
	Read    Permission = 1 << iota // 1 (001)
	Write                          // 2 (010)
	Execute                        // 4 (100)
)

// Typed vs untyped constant behaviour
const untypedInt = 100    // untyped — can be used as int, int64, float64 etc.
const typedInt int = 100  // typed — only usable as int

func main() {
	fmt.Println("=== Week 2: Constants (from go.dev/ref/spec) ===\n")

	fmt.Println("-- Typed & untyped constants --")
	fmt.Println("MaxTokens:", MaxTokens)
	fmt.Println("Pi:", Pi)

	// Untyped constant used in different numeric contexts
	var f64 float64 = untypedInt  // works — untyped adapts
	var i64 int64   = untypedInt  // works — untyped adapts
	fmt.Printf("untypedInt as float64: %f, as int64: %d\n", f64, i64)

	fmt.Println("\n-- iota: Weekday enum --")
	for d := Sunday; d <= Saturday; d++ {
		fmt.Printf("  %d = %s\n", d, d)
	}

	fmt.Println("\n-- iota: Permission bitmask --")
	fmt.Printf("  Read    = %03b (%d)\n", Read, Read)
	fmt.Printf("  Write   = %03b (%d)\n", Write, Write)
	fmt.Printf("  Execute = %03b (%d)\n", Execute, Execute)

	// Combine permissions using bitwise OR
	rw := Read | Write
	fmt.Printf("  Read|Write = %03b (%d)\n", rw, rw)

	// Check a permission with bitwise AND
	if rw&Read != 0 {
		fmt.Println("  Has Read permission: true")
	}
	if rw&Execute == 0 {
		fmt.Println("  Has Execute permission: false")
	}
}
