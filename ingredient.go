package main

type QuantityAmount float64

type Quantity struct {
	Amount QuantityAmount
	Unit   string
}

type Ingredient struct {
	Name     string
	Quantity Quantity
}

type IngredientList interface {
	IngredientsList() []Ingredient
}
