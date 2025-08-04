package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	ErrInvalidFormat = errors.New("invalid_format")
	ErrWordExists    = errors.New("word_exists")
)

func NewDefaultDictionary() map[string]string {
	return map[string]string{
		"makan":        "Eat",
		"minum":        "Drink",
		"tas punggung": "Backpack",
		"tertawa":      "Laugh",
		"tidur":        "Sleep",
	}
}

func Translate(dict map[string]string, word string) (string, bool) {
	id := strings.ToLower(strings.TrimSpace(word))
	if id == "" {
		return "", false
	}

	en, ok := dict[id]
	return en, ok
}

func AddWord(dict map[string]string, raw string) error {
	parts := strings.SplitN(raw, "#", 2)
	if len(parts) != 2 {
		return ErrInvalidFormat
	}

	idPart := strings.TrimSpace(parts[0])
	enPart := strings.TrimSpace(parts[1])

	if idPart == "" || enPart == "" {
		return ErrInvalidFormat
	}

	idKey := strings.ToLower(idPart)

	if _, exists := dict[idKey]; exists {
		return ErrWordExists
	}

	dict[idKey] = enPart
	return nil
}

func RemoveWord(dict map[string]string, word string) bool {
	id := strings.ToLower(strings.TrimSpace(word))
	if id == "" {
		return false
	}

	if _, ok := dict[id]; !ok {
		return false
	}

	delete(dict, id)
	return true
}

func BuildDictionaryLines(dict map[string]string) []string {
	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	lines := make([]string, 0, len(keys))
	for i, k := range keys {
		lines = append(lines, fmt.Sprintf("%d. %s: %s", i+1, k, dict[k]))
	}

	return lines
}

func printMenu() {
	fmt.Println("ID to EN Dictionary")
	fmt.Println("Menu:")
	fmt.Println("1. Translate")
	fmt.Println("2. Add word")
	fmt.Println("3. Remove word")
	fmt.Println("4. Print dictionary")
}

func main() {
	dict := NewDefaultDictionary()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		fmt.Print("\nInput: ")
		if !scanner.Scan() {
			return
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Word to translate: ")
			if !scanner.Scan() {
				return
			}
			word := scanner.Text()

			en, ok := Translate(dict, word)
			if !ok {
				fmt.Printf("\nsorry, %q is not found in dictionary\n\n", word)
				continue
			}

			fmt.Printf("\nID: %s\nEN: %s\n\n", word, en)

		case "2":
			fmt.Print("Word to be added in dict: ")
			if !scanner.Scan() {
				return
			}
			raw := scanner.Text()

			err := AddWord(dict, raw)
			if err != nil {
				if errors.Is(err, ErrInvalidFormat) {
					fmt.Println("\nword must be separated with # char")
				} else if errors.Is(err, ErrWordExists) {
					parts := strings.SplitN(raw, "#", 2)
					idPart := strings.TrimSpace(parts[0])
					fmt.Printf("\ncannot add existing word %q\n\n", strings.ToLower(idPart))
				}
				continue
			}

			fmt.Println("\nnew word successfully added")

		case "3":
			fmt.Print("Word to be removed: ")
			if !scanner.Scan() {
				return
			}
			word := scanner.Text()

			ok := RemoveWord(dict, word)
			if !ok {
				fmt.Printf("\nsorry, %q is not found in dictionary\n\n", word)
				continue
			}

			fmt.Printf("\n%q has been removed\n\n", word)

		case "4":
			fmt.Println()
			fmt.Println("ID to EN Dictionary:")
			lines := BuildDictionaryLines(dict)
			for _, line := range lines {
				fmt.Println(line)
			}
			fmt.Println()

		default:
			return
		}
	}
}
