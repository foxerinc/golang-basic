package main

import "testing"

func TestBuildZeroSquare_N2(t *testing.T) {
	got := BuildZeroSquare(2)
	want := "00\n00"

	if got != want {
		t.Errorf("BuildZeroSquare(2) = %q, want %q", got, want)
	}
}

func TestBuildZeroSquare_N3(t *testing.T) {
	got := BuildZeroSquare(3)
	want := "000\n000\n000"

	if got != want {
		t.Errorf("BuildZeroSquare(3) = %q, want %q", got, want)
	}
}

func TestBuildZeroSquare_NonPositive(t *testing.T) {
	got := BuildZeroSquare(0)
	want := ""

	if got != want {
		t.Errorf("BuildZeroSquare(0) = %q, want %q", got, want)
	}
}

func TestSumFirstN_Basic(t *testing.T) {
	got := SumFirstN(5)
	want := 15

	if got != want {
		t.Errorf("SumFirstN(5) = %d, want %d", got, want)
	}
}

func TestSumFirstN_Zero(t *testing.T) {
	got := SumFirstN(0)
	want := 0

	if got != want {
		t.Errorf("SumFirstN(0) = %d, want %d", got, want)
	}
}

func TestBuildMultiplicationTable_5(t *testing.T) {
	lines := BuildMultiplicationTable(5)

	if len(lines) != 10 {
		t.Fatalf("len(BuildMultiplicationTable(5)) = %d, want 10", len(lines))
	}

	if lines[0] != "5 x 1 = 5" {
		t.Errorf("lines[0] = %q, want %q", lines[0], "5 x 1 = 5")
	}

	if lines[9] != "5 x 10 = 50" {
		t.Errorf("lines[9] = %q, want %q", lines[9], "5 x 10 = 50")
	}
}

func TestReverseDigits_Positive(t *testing.T) {
	got := ReverseDigits(1234)
	want := 4321

	if got != want {
		t.Errorf("ReverseDigits(1234) = %d, want %d", got, want)
	}
}

func TestReverseDigits_WithTrailingZeros(t *testing.T) {
	got := ReverseDigits(1200)
	want := 21

	if got != want {
		t.Errorf("ReverseDigits(1200) = %d, want %d", got, want)
	}
}

func TestReverseDigits_Zero(t *testing.T) {
	got := ReverseDigits(0)
	want := 0

	if got != want {
		t.Errorf("ReverseDigits(0) = %d, want %d", got, want)
	}
}

func TestReverseDigits_Negative(t *testing.T) {
	got := ReverseDigits(-123)
	want := -321

	if got != want {
		t.Errorf("ReverseDigits(-123) = %d, want %d", got, want)
	}
}
