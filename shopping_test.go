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
					Quantity: Quantity{
						Amount: 1,
						Unit:   "Box",
					},
				},
				Ingredient{
					Name: "Canned Tomatoes",
					Quantity: Quantity{
						Amount: 1,
						Unit:   "Can",
					},
				},
				Ingredient{
					Name: "Onion",
					Quantity: Quantity{
						Amount: 1,
					},
				},
			},
		},
		Recipe{
			Ingredients: []Ingredient{
				Ingredient{
					Name: "Liver",
					Quantity: Quantity{
						Amount: 1,
					},
				},
				Ingredient{
					Name: "Onion",
					Quantity: Quantity{
						Amount: 1,
					},
				},
			},
		},
	}

	list := ShoppingList(recipes, nil)
	correctCount := 4
	if len(list) != correctCount {
		t.Fatalf("Expected %d items, got %d", correctCount, len(list))
	}
	assertStrEqual(t, "1 Box Penne Rigate", list[0])
	assertStrEqual(t, "2 Onion", list[2])
}

func TestShoppingListWithPantry(t *testing.T) {
	recipes := []Recipe{
		Recipe{
			Ingredients: []Ingredient{
				Ingredient{
					Name: "Onion",
					Quantity: Quantity{
						Amount: 1,
					},
				},
			},
		},
		Recipe{
			Ingredients: []Ingredient{
				Ingredient{
					Name: "Liver",
					Quantity: Quantity{
						Amount: 1,
					},
				},
				Ingredient{
					Name: "Onion",
					Quantity: Quantity{
						Amount: 1,
					},
				},
			},
		},
	}

	pantry := Pantry{
		ingredients: []Ingredient{
			Ingredient{
				Name: "Onion",
				Quantity: Quantity{
					Amount: 1,
				},
			},
		},
	}

	list := ShoppingList(recipes, &pantry)
	correctCount := 2
	if len(list) != correctCount {
		t.Fatalf("Expected %d items, got %d", correctCount, len(list))
	}
	assertStrEqual(t, "1 Onion", list[0])
}
