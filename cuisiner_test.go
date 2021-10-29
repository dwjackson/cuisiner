package main

import "testing"

func TestIngredientLineWithAmountOneAndNoUnit(t *testing.T) {
	ingredient := fakeIngredient("onion", 1, "")
	line := createIngredientLine(&ingredient)
	assertStrEqual(t, "onion", line)
}

func fakeIngredient(name string, amount QuantityAmount, unit string) Ingredient {
	return Ingredient{
		Name: name,
		Quantity: Quantity{
			Amount: amount,
			Unit:   unit,
		},
	}
}

func TestIngredientLineWithAmountOneAndUnit(t *testing.T) {
	ingredient := fakeIngredient("olive oil", 1, "tbsp")
	line := createIngredientLine(&ingredient)
	assertStrEqual(t, "1 tbsp olive oil", line)
}

func TestIngredientWithAmountAndNoUnit(t *testing.T) {
	ingredient := fakeIngredient("potatoes", 3, "")
	line := createIngredientLine(&ingredient)
	assertStrEqual(t, "3 potatoes", line)
}

func TestIngredientWithFractionalAmountAndNoUnit(t *testing.T) {
	ingredient := fakeIngredient("onions", 1.5, "")
	line := createIngredientLine(&ingredient)
	assertStrEqual(t, "1 ½ onions", line)
}

func TestIngredientWithFractionalAmountAndUnit(t *testing.T) {
	ingredient := fakeIngredient("hot pepper flakes", 1.5, "tsp")
	line := createIngredientLine(&ingredient)
	assertStrEqual(t, "1 ½ tsp hot pepper flakes", line)
}

func TestIngredientWithFractionalAmountLessThanOneAndUnit(t *testing.T) {
	ingredient := fakeIngredient("salt", 0.5, "tsp")
	line := createIngredientLine(&ingredient)
	assertStrEqual(t, "½ tsp salt", line)
}
