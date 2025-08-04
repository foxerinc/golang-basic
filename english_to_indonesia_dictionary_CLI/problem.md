# English to Indonesia Dictionary CLI

## Objectives

- Use the `map` data type to store dictionary entries.
- Use `for-range` to iterate over maps and slices.
- Handle basic error cases (word not found, invalid input, duplicate entries).
- Use `sort` from the Go standard library to print the dictionary in order.

## Description

You will build a simple CLI-based dictionary application that translates **Indonesian (ID)** words to **English (EN)**.

The program shows a menu:

```text
ID to EN Dictionary
Menu:
1. Translate
2. Add word
3. Remove word
4. Print dictionary
```

Behind the scenes, the dictionary is stored as a `map[string]string` where:

- The **key** is the Indonesian word in lowercase (e.g. `"makan"`).
- The **value** is the English translation (e.g. `"Eat"`).

The program should treat lookups and removals as **case-insensitive** by lowercasing user input.

## Initial Dictionary

The default dictionary (created in `NewDefaultDictionary`) contains:

- `makan` → `Eat`
- `minum` → `Drink`
- `tas punggung` → `Backpack`
- `tertawa` → `Laugh`
- `tidur` → `Sleep`

## Features

### 1. Translate an Indonesian word (Menu 1)

- Prompt:

  ```text
  Word to translate: <word>
  ```

- Behavior:
  - Convert the input word to lowercase and look it up in the map.
  - If found:

    ```text
    ID: <original input>
    EN: <translation>
    ```

  - If not found:

    ```text
    sorry, "<word>" is not found in dictionary
    ```

- Implemented by:

  ```go
  func Translate(dict map[string]string, word string) (string, bool)
  ```

  Returns the English translation and a boolean indicating whether it was found.

### 2. Add a new word (Menu 2)

- Prompt:

  ```text
  Word to be added in dict: <ID>#<EN>
  ```

- Rules:
  - Must use `#` as a separator between the Indonesian word and English word.
  - Both sides are trimmed from spaces.
  - Indonesian word is stored in lowercase.
  - If the separator is missing or either part is empty:

    ```text
    word must be separated with # char
    ```

  - If the word already exists in the dictionary (case-insensitive):

    ```text
    cannot add existing word "<id>"
    ```

  - If the word is valid and new:

    ```text
    new word successfully added
    ```

- Implemented by:

  ```go
  func AddWord(dict map[string]string, raw string) error
  ```

  This function parses the raw input string and:

  - Returns `ErrInvalidFormat` if the separator or format is wrong.
  - Returns `ErrWordExists` if the ID word already exists.
  - Returns `nil` if the word is added successfully.

### 3. Remove a word (Menu 3)

- Prompt:

  ```text
  Word to be removed: <word>
  ```

- Behavior:
  - Convert input to lowercase and check if it exists in the dictionary.
  - If found:

    ```text
    "<word>" has been removed
    ```

  - If not found:

    ```text
    sorry, "<word>" is not found in dictionary
    ```

- Implemented by:

  ```go
  func RemoveWord(dict map[string]string, word string) bool
  ```

  Returns `true` if the word existed and was removed, `false` otherwise.

### 4. Print dictionary (Menu 4)

- Output example:

  ```text
  ID to EN Dictionary:
  1. membaca: Read
  2. minum: Drink
  3. tas punggung: Backpack
  4. tertawa: Laugh
  5. tidur: Sleep
  ```

- Requirements:
  - Print entries sorted **ascending by Indonesian word**.
  - Use `sort.Strings` from the standard library.

- Implemented helper:

  ```go
  func BuildDictionaryLines(dict map[string]string) []string
  ```

  Which returns a slice of strings like:

  ```text
  1. membaca: Read
  2. minum: Drink
  ...
  ```

The CLI code uses this helper to print the lines.

## Core Helper Functions

### NewDefaultDictionary

```go
func NewDefaultDictionary() map[string]string
```

Returns a map with the initial words:

- `makan` → `Eat`
- `minum` → `Drink`
- `tas punggung` → `Backpack`
- `tertawa` → `Laugh`
- `tidur` → `Sleep`

### Translate

```go
func Translate(dict map[string]string, word string) (string, bool)
```

- Lowercases the input.
- Returns the English translation and `true` if found.
- Returns `""` and `false` if not found.

### AddWord

```go
func AddWord(dict map[string]string, raw string) error
```

- Parses input in the form `"ID#EN"`.
- Uses `#` as separator.
- Lowercases and trims the ID.
- Trims the English word.
- Returns:

  - `ErrInvalidFormat` if the format is wrong.
  - `ErrWordExists` if the word already exists.
  - `nil` on success.

### RemoveWord

```go
func RemoveWord(dict map[string]string, word string) bool
```

- Lowercases and trims the input.
- Deletes the entry from the map if it exists.
- Returns `true` if deleted, `false` otherwise.

### BuildDictionaryLines

```go
func BuildDictionaryLines(dict map[string]string) []string
```

- Collects all keys into a slice.
- Sorts the keys using `sort.Strings`.
- Builds lines with format:

  ```text
  <index>. <id>: <en>
  ```

- Returns the slice of lines.

## Example Interactions

### Case: word found in dictionary

```text
ID to EN Dictionary
Menu:
1. Translate
2. Add word
3. Remove word
4. Print dictionary

Input: 1
Word to translate: Makan

ID: Makan
EN: Eat
```

### Case: word found (case insensitive)

```text
Input: 1
Word to translate: makan

ID: makan
EN: Eat
```

### Case: word not registered

```text
Input: 1
Word to translate: Terima kasih

sorry, "Terima kasih" is not found in dictionary
```

### Case: add new word

```text
Input: 2
Word to be added in dict: Membaca#Read

new word successfully added
```

### Case: add existing word

```text
Input: 2
Word to be added in dict: makan#eat

cannot add existing word "makan"
```

### Case: wrong separator

```text
Input: 2
Word to be added in dict: jalan-walk

word must be separated with # char
```

### Case: remove existing word

```text
Input: 3
Word to be removed: makan

"makan" has been removed
```

### Case: remove missing word

```text
Input: 3
Word to be removed: Lato-lato

sorry, "Lato-lato" is not found in dictionary
```

### Case: print dictionary

After adding `membaca` and removing `makan`:

```text
Input: 4

ID to EN Dictionary:
1. membaca: Read
2. minum: Drink
3. tas punggung: Backpack
4. tertawa: Laugh
5. tidur: Sleep
```

## How to Run

Run the CLI:

```bash
go run .
```

Run the tests:

```bash
go test ./...
```