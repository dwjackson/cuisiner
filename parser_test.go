package cuisiner

import (
	"testing"
)

func TestParsePlainText(t *testing.T) {
	input := "Flip the pancake."
	recipe, err := Parse(input)
	if err != nil || len(recipe.Directions) != 1 {
		t.Fatalf("Expected 1 direction")
	} else {
		direction := recipe.Directions[0]
		testDirection(t, input, direction)
	}
}

func TestParseMultipleLinesOfPlainText(t *testing.T) {
	input := "Flip the pancake.\nFlip the pancake again."
	recipe, err := Parse(input)
	if err != nil || len(recipe.Directions) != 2 {
		t.Fatalf("Expected 2 directions")
	}
}

func TestSkipBlankLines(t *testing.T) {
	input := "Flip the pancake.\n \t\nFlip the pancake again."
	recipe, err := Parse(input)
	if err != nil || len(recipe.Directions) != 2 {
		t.Fatalf("Expected 2 directions")
	}
}

func testDirection(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Fatalf("Expected %q, got %q", expected, actual)
	}
}
