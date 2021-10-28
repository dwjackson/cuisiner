package main

import (
	"testing"
)

func TestShoppingList(t *testing.T) {
	recipes := []Recipe{
		Recipe{
			Ingredients: []Ingredient{
				Ingredient{
					Name: "Penne Rigate",
					Quantity: 1,
					Unit: "Box",
				},
				Ingredient{
					Name: "Canned Tomatoes",
					Quantity: 1,
					Unit: "Can",
				},
				Ingredient{
					Name: "Onion",
					Quantity: 1,
				},
			},
		},
		Recipe{
			Ingredients: []Ingredient{
				Ingredient{
					Name: "Liver",
					Quantity: 1,
				},
				Ingredient{
					Name: "Onion",
					Quantity: 1,
				},
			},
		},
	}

	list := ShoppingList(recipes)
	correctCount := 4
	if len(list) != correctCount {
		t.Fatalf("Expected %d items, got %d", correctCount, len(list))
	}
	// TODO
}
