package main

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

func TestIngredientWithoutQuantityWithUnicodeInName(t *testing.T) {
	input := "Add the @gruyère"
	recipe, err := Parse(input)
	if err != nil || len(recipe.Directions) != 1 {
		t.Fatalf("Expected 1 direction")
	}
	if len(recipe.Ingredients) != 1 {
		t.Fatalf("Expected 1 ingredient")
	}
	assertStrEqual(t, "gruyère", recipe.Ingredients[0].Name)
	assertStrEqual(t, "Add the gruyère", recipe.Directions[0])
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

func TestIngredientWithSpaceInName(t *testing.T) {
	input := "Put the @crème fraîche{} into the bowl"
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}
	if len(recipe.Ingredients) != 1 {
		t.Fatalf("Expected 1 ingredient")
	} else {
		ingredient := recipe.Ingredients[0]
		assertStrEqual(t, "crème fraîche", ingredient.Name)
		assertQuantityEqual(t, 1, ingredient.Quantity.Amount)

		direction := recipe.Directions[0]
		assertStrEqual(t, "Put the crème fraîche into the bowl", direction)
	}
}

func TestIngredientWithQuantity(t *testing.T) {
	input := "Chop up @potatos{2} and set aside"
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}
	if len(recipe.Ingredients) != 1 {
		t.Fatalf("Expected 1 ingredient")
	} else {
		ingredient := recipe.Ingredients[0]
		assertStrEqual(t, "potatos", ingredient.Name)
		assertQuantityEqual(t, 2, ingredient.Quantity.Amount)
	}
}

func TestIngredientsWithAndWithoutQuantity(t *testing.T) {
	input := "Chop up @potatos{2} and @leek and set aside"
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}
	if len(recipe.Ingredients) != 2 {
		t.Fatalf("Expected 2 ingredients")
	} else {
		potato := recipe.Ingredients[0]
		assertStrEqual(t, "potatos", potato.Name)
		assertQuantityEqual(t, 2, potato.Quantity.Amount)

		leek := recipe.Ingredients[1]
		assertStrEqual(t, "leek", leek.Name)
		assertQuantityEqual(t, 1, leek.Quantity.Amount)
	}
	direction := recipe.Directions[0]
	assertStrEqual(t, "Chop up potatos and leek and set aside", direction)
}

func TestIngredientsWithQuantityAndUnit(t *testing.T) {
	input := "Mix @water{300%mL} and @flour{400%g} in a bowl"
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}
	ingredientCount := len(recipe.Ingredients)
	if ingredientCount != 2 {
		t.Fatalf("Expected 2 ingredients, got %d", ingredientCount)
	} else {
		water := recipe.Ingredients[0]
		assertStrEqual(t, "water", water.Name)
		assertQuantityEqual(t, 300, water.Quantity.Amount)
		assertStrEqual(t, "mL", water.Quantity.Unit)

		flour := recipe.Ingredients[1]
		assertStrEqual(t, "flour", flour.Name)
		assertQuantityEqual(t, 400, flour.Quantity.Amount)
		assertStrEqual(t, "g", flour.Quantity.Unit)
	}
}

func TestIngredientWithFractionalQuantityAndUnit(t *testing.T) {
	input := "Add @sugar{2.5%tsp} to the bowl"
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}
	sugar := recipe.Ingredients[0]
	assertStrEqual(t, "sugar", sugar.Name)
	assertQuantityEqual(t, 2.5, sugar.Quantity.Amount)
	assertStrEqual(t, "tsp", sugar.Quantity.Unit)
}

func TestIngredientWithFractionalQuantityLessThanOneAndUnit(t *testing.T) {
	input := "Add @salt{0.5%tsp} to the bowl"
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}
	salt := recipe.Ingredients[0]
	assertStrEqual(t, "salt", salt.Name)
	assertQuantityEqual(t, 0.5, salt.Quantity.Amount)
	assertStrEqual(t, "tsp", salt.Quantity.Unit)
}

func TestSkipComment(t *testing.T) {
	input := "Add @sugar to bowl // add some sugar"
	recipe, _ := Parse(input)
	direction := recipe.Directions[0]
	assertStrEqual(t, "Add sugar to bowl", direction)
}

func TestTimer(t *testing.T) {
	input := "Bake for ~{15%minutes}."
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}
	assertIntEqual(t, 1, len(recipe.Timers))

	timer := recipe.Timers[0]
	assertDurationEqual(t, 15, timer.duration)
	assertStrEqual(t, "minutes", timer.Unit)

	direction := recipe.Directions[0]
	assertStrEqual(t, "Bake for 15 minutes.", direction)
}

func TestCookware(t *testing.T) {
	input := "Fill a #pot with water\nPut onions in the #sauce pan{}"
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}

	assertIntEqual(t, 2, len(recipe.Cookware))
	assertStrEqual(t, "pot", recipe.Cookware[0])
	assertStrEqual(t, "sauce pan", recipe.Cookware[1])
	assertStrEqual(t, "Fill a pot with water", recipe.Directions[0])
	assertStrEqual(t, "Put onions in the sauce pan", recipe.Directions[1])
}

func TestIngredientAndCookwareOnTheSameLine(t *testing.T) {
	input := "Put the chopped @carrot into the #sauce pan{}"
	recipe, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed")
	}

	assertIntEqual(t, 1, len(recipe.Cookware))
	assertIntEqual(t, 1, len(recipe.Ingredients))
	assertStrEqual(t, "carrot", recipe.Ingredients[0].Name)
	assertStrEqual(t, "sauce pan", recipe.Cookware[0])
}

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

func assertQuantityEqual(t *testing.T, expected QuantityAmount, actual QuantityAmount) {
	if actual != expected {
		t.Fatalf("Expected %f, got %f", expected, actual)
	}
}

func assertDurationEqual(t *testing.T, expected float64, actual float64) {
	if actual != expected {
		t.Fatalf("Expected %f, got %f", expected, actual)
	}
}
