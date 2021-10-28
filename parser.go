package cuisiner

import (
	"regexp"
	"strconv"
	"strings"
)

func Parse(input string) (*Recipe, error) {
	var ingredients []Ingredient
	var directions []string
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			// Skip empty lines
			continue
		}
		lineIngredients, parsedLine := discoverIngredients(line)
		for _, ingredient := range lineIngredients {
			ingredients = append(ingredients, ingredient)
		}
		directions = append(directions, parsedLine)
	}
	recipe := &Recipe{
		Ingredients: ingredients,
		Directions:  directions,
	}
	return recipe, nil
}

func discoverIngredients(line string) ([]Ingredient, string) {
	var ingredients []Ingredient

	reQuantity := regexp.MustCompile(`\@([^\@]+)\{(\d+)(\%[^\}]+)?\}`)
	for _, m := range reQuantity.FindAllStringSubmatch(line, -1) {
		name := m[1]
		quantity64, err := strconv.ParseInt(m[2], 10, 32)
		unitPart := m[3]
		var unit string
		if len(unitPart) > 0 {
			unit = unitPart[1:]
		} else {
			unit = ""
		}
		if err != nil {
			panic("Bad quantity") // TODO: Do not panic
		}
		quantity := int(quantity64)
		ingredient := Ingredient{
			Name:     name,
			Quantity: quantity,
			Unit: unit,
		}
		ingredients = append(ingredients, ingredient)
	}

	line = reQuantity.ReplaceAllString(line, "$1")

	reNoQuantity := regexp.MustCompile(`\@(\w+)`)
	for _, m := range reNoQuantity.FindAllString(line, -1) {
		name := m[1:]
		quantity := 1
		ingredient := Ingredient{
			Name:     name,
			Quantity: quantity,
		}
		ingredients = append(ingredients, ingredient)
	}

	line = reNoQuantity.ReplaceAllString(line, "$1")

	return ingredients, line
}
