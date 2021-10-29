package main

import (
	"testing"
)

func TestFormatIntegerAmount(t *testing.T) {
	ingredient := Ingredient{
		Name: "onion",
		Quantity: Quantity{
			Amount: 2,
		},
	}
	assertStrEqual(t, "2", ingredient.FormatAmount())
}

func TestFormatFractionalAmount(t *testing.T) {
	ingredient := Ingredient{
		Name: "onion",
		Quantity: Quantity{
			Amount: 2.5,
		},
	}
	assertStrEqual(t, "2 ½", ingredient.FormatAmount())
}

func TestFormatFractions(t *testing.T) {
	fractions := map[float64]string{
		0.125: "⅛",
		0.25:  "¼",
		0.5:   "½",
		0.75:  "¾",
	}
	for decimal, fraction := range fractions {
		ingredient := Ingredient{
			Name: "onion",
			Quantity: Quantity{
				Amount: QuantityAmount(decimal),
			},
		}
		assertStrEqual(t, fraction, ingredient.FormatAmount())
	}
}
