package main

import "fmt"

func main() {
	fmt.Println("=== Week 2: Go Literals (from go.dev/ref/spec) ===\n")

	// --- Integer literals ---
	decimal := 1_000_000    // underscore separator for readability
	binary := 0b1010_1111   // binary prefix 0b
	octal := 0o755          // octal prefix 0o (Unix permissions)
	hex := 0xFF_A0_3C       // hexadecimal prefix 0x

	fmt.Println("-- Integer literals --")
	fmt.Printf("decimal  : %d\n", decimal)
	fmt.Printf("binary   : %d (0b%b)\n", binary, binary)
	fmt.Printf("octal    : %d (0o%o)\n", octal, octal)
	fmt.Printf("hex      : %d (0x%X)\n", hex, hex)

	// --- Floating-point literals ---
	fmt.Println("\n-- Floating-point literals --")
	f1 := 3.14
	f2 := .5       // leading dot is valid
	f3 := 1.5e10   // scientific notation
	f4 := 1.5e-3
	fmt.Printf("3.14      = %f\n", f1)
	fmt.Printf(".5        = %f\n", f2)
	fmt.Printf("1.5e10    = %e\n", f3)
	fmt.Printf("1.5e-3    = %f\n", f4)

	// --- Imaginary literals ---
	fmt.Println("\n-- Imaginary literals --")
	c1 := 3 + 4i
	c2 := complex(1.5, -2.5)
	fmt.Printf("3+4i          = %v\n", c1)
	fmt.Printf("complex(1.5, -2.5) = %v\n", c2)
	fmt.Printf("magnitude of 3+4i  = %.2f\n", real(c1*c1)+imag(c1*c1)) // not actual magnitude, just demo

	// --- Rune literals ---
	fmt.Println("\n-- Rune literals --")
	r1 := 'A'           // simple character
	r2 := '\n'          // escape: newline
	r3 := '\u0041'      // unicode: 'A'
	r4 := '\U00000041'  // full unicode: 'A'
	fmt.Printf("'A'         = %c (code: %d)\n", r1, r1)
	fmt.Printf("'\\n'        = newline (code: %d)\n", r2)
	fmt.Printf("'\\u0041'    = %c (code: %d)\n", r3, r3)
	fmt.Printf("'\\U00000041'= %c (code: %d)\n", r4, r4)

	// --- String literals ---
	fmt.Println("\n-- String literals --")
	interpreted := "Hello\tGo\nSpec!"         // escape sequences processed
	raw := `Hello\tGo\nSpec!`                 // raw — no escape processing
	unicode := "नमस्ते"                        // UTF-8 encoded string

	fmt.Printf("interpreted : %s\n", interpreted)
	fmt.Printf("raw         : %s\n", raw)
	fmt.Printf("unicode     : %s (bytes: %d, runes: %d)\n",
		unicode, len(unicode), len([]rune(unicode)))
}
