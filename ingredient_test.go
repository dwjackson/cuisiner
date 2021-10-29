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
