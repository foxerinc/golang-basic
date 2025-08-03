package main

import (
	"fmt"
	"strings"
)

// Exercise 0: build a square of '0' characters with size N×N.
// If N <= 0, returns an empty string.
func BuildZeroSquare(n int) string {
	if n <= 0 {
		return ""
	}

	line := strings.Repeat("0", n)

	var b strings.Builder
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteByte('\n')
		}
		b.WriteString(line)
	}

	return b.String()
}

// Exercise 1: sum of the first N natural numbers (1..N).
// If N <= 0, returns 0.
func SumFirstN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// Exercise 2: multiplication table for N up to 10.
// Returns a slice of lines like "5 x 1 = 5".
func BuildMultiplicationTable(n int) []string {
	lines := make([]string, 0, 10)
	for i := 1; i <= 10; i++ {
		lines = append(lines, fmt.Sprintf("%d x %d = %d", n, i, n*i))
	}
	return lines
}

// Exercise 3: reverse digits of an integer.
// Handles negative numbers by keeping the sign.
func ReverseDigits(n int) int {
	if n == 0 {
		return 0
	}

	sign := 1
	if n < 0 {
		sign = -1
		n = -n
	}

	reversed := 0
	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}

	return sign * reversed
}

func main() {
	// Exercise 0 demo
	fmt.Println("Exercise 0: square of 0 (N = 3)")
	fmt.Println(BuildZeroSquare(3))
	fmt.Println()

	// Exercise 1 demo
	fmt.Println("Exercise 1: sum of first N natural numbers (N = 5)")
	n := 5
	sum := SumFirstN(n)
	fmt.Printf("1 + 2 + 3 + 4 + 5 = %d\n\n", sum)

	// Exercise 2 demo
	fmt.Println("Exercise 2: multiplication table (N = 5)")
	for _, line := range BuildMultiplicationTable(5) {
		fmt.Println(line)
	}
	fmt.Println()

	// Exercise 3 demo
	fmt.Println("Exercise 3: reverse digits (1234)")
	fmt.Printf("Reversed: %d\n", ReverseDigits(1234))
}
