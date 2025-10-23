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

var knownFractions = map[float64]string{
	0.125: "⅛",
	0.25:  "¼",
	0.375: "⅜",
	0.5:   "½",
	0.625: "⅝",
	0.75:  "¾",
	0.875: "⅞",
}

func (i *Ingredient) FormatAmount() string {
	amount := float64(i.Quantity.Amount)

	isInteger := math.Floor(amount) == amount
	if isInteger {
		return strconv.Itoa(int(amount))
	}

	wholePart := math.Floor(amount)
	fractionPart := amount - wholePart
	if fraction, isFraction := knownFractions[fractionPart]; isFraction {
		if wholePart > 0 {
			return fmt.Sprintf("%d %s", int(wholePart), fraction)
		}
		return fraction
	}

	amountString := fmt.Sprintf("%5.5f", amount)
	amountString = strings.TrimRight(amountString, "0")
	return amountString
}

type IngredientList interface {
	IngredientsList() []Ingredient
}
