package main

import "strconv"

func ShoppingList(recipes []Recipe) []string {
	var itemOrder []string
	itemQuantities := make(map[string]float64)

	for _, recipe := range recipes {
		for _, ingredient := range recipe.Ingredients {
			unit := ingredient.Quantity.Unit
			var item string
			if len(unit) > 0 {
				item = ingredient.Quantity.Unit + " " + ingredient.Name
			} else {
				item = ingredient.Name
			}
			if _, exists := itemQuantities[item]; !exists {
				itemQuantities[item] = 0.0
				itemOrder = append(itemOrder, item)
			}
			itemQuantities[item] += float64(ingredient.Quantity.Amount)
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
