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

func TestIngredientWithoutQuantity(t *testing.T) {
	input := "Chop up a @potato and set aside"
	recipe, err := Parse(input)
	if err != nil || len(recipe.Directions) != 1 {
		t.Fatalf("Expected 1 direction")
	}
	if len(recipe.Ingredients) != 1 {
		t.Fatalf("Expected 1 ingredient")
	}
	if recipe.Ingredients[0] != "potato" {
		t.Fatalf("Expected ingredient potato, got %q", recipe.Ingredients[0])
	}
}

func TestTwoIngredientsWithoutQuantity(t *testing.T) {
	input := "Chop up a @potato and a @leek and set aside"
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}
	if len(recipe.Ingredients) != 2 {
		t.Fatalf("Expected 2 ingredients")
	}
	if recipe.Ingredients[0] != "potato" {
		t.Fatalf("Expected ingredient potato, got %q", recipe.Ingredients[0])
	}
	if recipe.Ingredients[1] != "leek" {
		t.Fatalf("Expected ingredient leek, got %q", recipe.Ingredients[1])
	}
}

func testDirection(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Fatalf("Expected %q, got %q", expected, actual)
	}
}
