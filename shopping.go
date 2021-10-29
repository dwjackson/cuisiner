package main

import "strconv"

type Pantry struct {
	ingredients []Ingredient
}

func (p *Pantry) IngredientsList() []Ingredient {
	return p.ingredients
}

func ShoppingList(recipes []Recipe, pantry *Pantry) []string {
	var itemOrder []string
	itemQuantities := make(map[string]float64)

	for _, recipe := range recipes {
		for _, ingredient := range recipe.IngredientsList() {
			item := itemName(&ingredient)
			if _, exists := itemQuantities[item]; !exists {
				itemQuantities[item] = 0.0
				itemOrder = append(itemOrder, item)
			}
			amount := ingredient.Quantity.Amount
			itemQuantities[item] += float64(amount)
		}
	}

	if pantry != nil {
		for _, ingredient := range pantry.IngredientsList() {
			item := itemName(&ingredient)
			if _, exists := itemQuantities[item]; exists {
				amount := float64(ingredient.Quantity.Amount)
				itemQuantities[item] -= amount
			}
		}
	}

	var list []string
	for _, item := range itemOrder {
		amount := itemQuantities[item]
		quantity := strconv.FormatFloat(amount, 'f', -1, 64)
		list = append(list, quantity+" "+item)
	}
	return list
}

func itemName(ingredient *Ingredient) string {
	unit := ingredient.Quantity.Unit
	var item string
	if len(unit) > 0 {
		item = ingredient.Quantity.Unit + " " + ingredient.Name
	} else {
		item = ingredient.Name
	}
	return item
}
