package main

type Recipe struct {
	Ingredients []Ingredient
	Directions  []string
	Timers      []Timer
	Cookware    []string
}

func (r *Recipe) IngredientsList() []Ingredient {
	return r.Ingredients
}
