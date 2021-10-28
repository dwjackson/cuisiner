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
		assertStrEqual(t, input, direction)
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
	assertStrEqual(t, "potato", recipe.Ingredients[0].Name)
}

func TestTwoIngredientsWithoutQuantity(t *testing.T) {
	input := "Chop up a @potato and a @leek and set aside"
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}
	if len(recipe.Ingredients) != 2 {
		t.Fatalf("Expected 2 ingredients")
	} else {
		assertStrEqual(t, "potato", recipe.Ingredients[0].Name)
		assertStrEqual(t, "leek", recipe.Ingredients[1].Name)
	}
}

func TestIngredientWithQuantity(t *testing.T) {
	input := "Chop up @potato{2} and set aside"
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}
	if len(recipe.Ingredients) != 1 {
		t.Fatalf("Expected 1 ingredient")
	} else {
		ingredient := recipe.Ingredients[0]
		assertStrEqual(t, "potato", ingredient.Name)
		assertIntEqual(t, 2, ingredient.Quantity)
	}
}

// TODO: Test ingredients with and without quantity in same line

func assertStrEqual(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Fatalf("Expected %q, got %q", expected, actual)
	}
}

func assertIntEqual(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
