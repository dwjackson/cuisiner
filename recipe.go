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

type Recipe struct {
	Ingredients []Ingredient
	Directions  []string
	Timers      []Timer
	Cookware    []string
}

type Timer struct {
	Duration float64
	Unit     string
}
