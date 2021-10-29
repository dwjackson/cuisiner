package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type QuantityAmount float64

type Quantity struct {
	Amount QuantityAmount
	Unit   string
}

type Ingredient struct {
	Name     string
	Quantity Quantity
}

func (i *Ingredient) FormatAmount() string {
	amount := float64(i.Quantity.Amount)

	isInteger := math.Floor(amount) == amount
	if isInteger {
		return strconv.Itoa(int(amount))
	}

	amountString := fmt.Sprintf("%5.5f", amount)
	amountString = strings.TrimRight(amountString, "0")
	return amountString
}

type IngredientList interface {
	IngredientsList() []Ingredient
}
