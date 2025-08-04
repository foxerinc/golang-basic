package main

import (
	"errors"
	"testing"
)

func TestTranslate_Found_CaseInsensitive(t *testing.T) {
	dict := NewDefaultDictionary()

	en, ok := Translate(dict, "Makan")
	if !ok {
		t.Fatalf("expected word to be found")
	}
	if en != "Eat" {
		t.Errorf("Translate(Makan) = %q, want %q", en, "Eat")
	}

	en2, ok2 := Translate(dict, "makan")
	if !ok2 {
		t.Fatalf("expected word to be found for lowercase")
	}
	if en2 != "Eat" {
		t.Errorf("Translate(makan) = %q, want %q", en2, "Eat")
	}
}

func TestTranslate_NotFound(t *testing.T) {
	dict := NewDefaultDictionary()

	en, ok := Translate(dict, "Terima kasih")
	if ok {
		t.Fatalf("expected not found, got ok with %q", en)
	}
}

func TestAddWord_Success(t *testing.T) {
	dict := make(map[string]string)

	err := AddWord(dict, "Membaca#Read")
	if err != nil {
		t.Fatalf("unexpected error adding word: %v", err)
	}

	if dict["membaca"] != "Read" {
		t.Errorf("dict[\"membaca\"] = %q, want %q", dict["membaca"], "Read")
	}
}

func TestAddWord_Existing(t *testing.T) {
	dict := NewDefaultDictionary()

	err := AddWord(dict, "makan#eat")
	if !errors.Is(err, ErrWordExists) {
		t.Fatalf("expected ErrWordExists, got %v", err)
	}
}

func TestAddWord_InvalidFormat_NoSeparator(t *testing.T) {
	dict := make(map[string]string)

	err := AddWord(dict, "jalan-walk")
	if !errors.Is(err, ErrInvalidFormat) {
		t.Fatalf("expected ErrInvalidFormat, got %v", err)
	}
}

func TestRemoveWord_Existing(t *testing.T) {
	dict := NewDefaultDictionary()

	ok := RemoveWord(dict, "makan")
	if !ok {
		t.Fatalf("expected remove to return true")
	}

	if _, exists := dict["makan"]; exists {
		t.Errorf("expected \"makan\" to be deleted from dict")
	}
}

func TestRemoveWord_NotFound(t *testing.T) {
	dict := NewDefaultDictionary()

	ok := RemoveWord(dict, "Lato-lato")
	if ok {
		t.Fatalf("expected remove to return false for non-existing word")
	}
}

func TestBuildDictionaryLines_Sorted(t *testing.T) {
	dict := map[string]string{
		"tas punggung": "Backpack",
		"minum":        "Drink",
		"membaca":      "Read",
		"tidur":        "Sleep",
		"tertawa":      "Laugh",
	}

	lines := BuildDictionaryLines(dict)

	if len(lines) != 5 {
		t.Fatalf("len(lines) = %d, want 5", len(lines))
	}

	wantFirst := "1. membaca: Read"
	if lines[0] != wantFirst {
		t.Errorf("lines[0] = %q, want %q", lines[0], wantFirst)
	}

	wantLast := "5. tidur: Sleep"
	if lines[4] != wantLast {
		t.Errorf("lines[4] = %q, want %q", lines[4], wantLast)
	}
}
