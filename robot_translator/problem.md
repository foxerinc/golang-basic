# Robot Translator

## Objectives
- Use conditional statements.
- Use for-loops.
- Convert a string into a slice of tokens.
- Concatenate strings to build structured output.

## Description
A robot test facility needs a program that converts a compact command string into human‑readable, step‑by‑step instructions.

The robot understands three commands:
- `R` — turn **right**
- `L` — turn **left**
- `A` — **advance** forward

Your task is to **group consecutive identical commands** and output one instruction line per group.

## Input
A single string of length ≥ 1 composed only of the characters `R`, `L`, and `A`.  
Example: `RRAAALA`

## Output
A **line-feed separated string** where each line has the form:
- `Move right <n> time(s)`
- `Move left <n> time(s)`
- `Move advance <n> time(s)`

Use pluralization: append `s` to `time` when `<n> > 1`.

## Example
**Input**  
```
RRAAALA
```

**Output**  
```
Move right 2 times
Move advance 3 times
Move left 1 time
Move advance 1 time
```

## Rules
1. Only `R`, `L`, and `A` are valid symbols; any other character makes the input invalid.
2. Group **consecutive** identical symbols.
3. Each group becomes one output line.
4. Print lines in the order the groups appear in the input.

## Suggested Approach
1. Convert the input string into a slice of single‑character strings or runes.
2. Iterate once through the slice while tracking the **current symbol** and a **counter**.
3. When the symbol changes, emit a line for the previous symbol and reset the counter.
4. After the loop, emit the final pending group.
5. Build the output using string concatenation or a `strings.Builder`.

## Complexity
- Time: `O(n)` where `n` is the length of the input string.
- Space: `O(1)` additional space beyond the output.

## How to run
```
go run .
```