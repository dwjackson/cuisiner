package cuisiner

import (
	"regexp"
	"strings"
)

type Ingredient struct {
	Name string
}

type Recipe struct {
	Ingredients []Ingredient
	Directions  []string
}

func Parse(input string) (*Recipe, error) {
	var ingredients []Ingredient
	var directions []string
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			lineIngredients := discoverIngredients(line)
			for _, ingredientName := range lineIngredients {
				ingredient := Ingredient{
					Name: ingredientName,
				}
				ingredients = append(ingredients, ingredient)
			}
			directions = append(directions, line)
		}
	}
	recipe := &Recipe{
		Ingredients: ingredients,
		Directions:  directions,
	}
	return recipe, nil
}

func discoverIngredients(line string) []string {
	var ingredients []string
	re := regexp.MustCompile(`\@\w+`)
	matches := re.FindAllString(line, -1)
	for _, m := range matches {
		name := m[1:]
		ingredients = append(ingredients, name)
	}
	return ingredients
}
