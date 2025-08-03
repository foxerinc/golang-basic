# Basic Go Exercises — Loops, Numbers, and Patterns

This folder contains four small Go exercises focused on loops, basic arithmetic, and simple string handling.

---

## Exercise 0 — Square of 0

### Objectives

- Practice working with loops and string manipulation.
- Understand how to construct patterns using nested loops.

### Description

Write a program to print a square of `0` given the input of `N`.

**Input**  
An integer `N` representing the size of the square.

**Output**  
A square pattern of `0` with dimensions `N x N`.

### Examples

For `N = 2`:

```text
00
00
```

For `N = 3`:

```text
000
000
000
```

### Implementation

Implemented as:

```go
func BuildZeroSquare(n int) string
```

- Returns a multiline string with `N` lines.
- Each line contains `N` characters of `0`.
- If `n <= 0`, returns an empty string.

---

## Exercise 1 — Sum of First N Natural Numbers

### Objectives

- Understand how to use loops to calculate sums.
- Practice basic arithmetic operations.

### Description

Write a program to calculate the sum of the first `N` natural numbers.

**Input**  
An integer `N`.

**Output**  
The sum of the first `N` natural numbers.

### Example

For `N = 5`:

```text
1 + 2 + 3 + 4 + 5 = 15
```

### Implementation

Implemented as:

```go
func SumFirstN(n int) int
```

- Uses a `for` loop from `1` to `n`.
- Accumulates the sum in a variable.
- Returns `0` if `n <= 0`.

---

## Exercise 2 — Multiplication Table

### Objectives

- Practice using loops to generate patterns of output.
- Understand how to create a multiplication table.

### Description

Write a program that prints the multiplication table for a given number `N` up to `10`.

**Input**  
An integer `N`.

**Output**  
The multiplication table for `N`.

### Example

For `N = 5`:

```text
5 x 1 = 5
5 x 2 = 10
5 x 3 = 15
...
5 x 10 = 50
```

### Implementation

Implemented as:

```go
func BuildMultiplicationTable(n int) []string
```

- Returns a slice of strings.
- Each element has the format: `"<N> x <i> = <result>"` for `i` from `1` to `10`.

---

## Exercise 3 — Reverse Digits of a Number

### Objectives

- Practice working with numbers and digit manipulation.
- Understand how to reverse the digits of a number.

### Description

Write a program to reverse the digits of a given number.

**Input**  
An integer.

**Output**  
The reversed integer.

### Example

For input `1234`:

```text
Output: 4321
```

### Implementation

Implemented as:

```go
func ReverseDigits(n int) int
```

- Extracts digits using `% 10` and `/ 10`.
- Builds the reversed number step by step.
- Handles negative numbers by keeping the sign.
- `ReverseDigits(0)` returns `0`.

---

## Main Program

The `main` function can be used to demonstrate:

- `BuildZeroSquare(3)`
- `SumFirstN(5)`
- `BuildMultiplicationTable(5)`
- `ReverseDigits(1234)`

and print outputs similar to the examples above.

### How to Run

Run the program:

```bash
go run .
```

Run the tests:

```bash
go test ./...
```