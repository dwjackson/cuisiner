package cuisiner

import (
	"strings"
	"regexp"
)

type Recipe struct {
	Ingredients []string
	Directions []string
}

func Parse(input string) (*Recipe, error) {
	var ingredients []string
	var directions []string
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			lineIngredients := discoverIngredients(line)
			for _, ingredient := range lineIngredients {
				ingredients = append(ingredients, ingredient)
			}
			directions = append(directions, line)
		}
	}
	recipe := &Recipe{
		Ingredients: ingredients,
		Directions: directions,
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
