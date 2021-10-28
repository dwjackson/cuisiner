package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Parse(input string) (*Recipe, error) {
	var ingredients []Ingredient
	var directions []string
	var timers []Timer
	var cookware []string
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := removeComments(line)
		if len(trimmedLine) == 0 {
			// Skip empty lines
			continue
		}

		lineIngredients, parsedLine, ingredientsError := discoverIngredients(trimmedLine)
		if ingredientsError != nil {
			return nil, ingredientsError
		}
		for _, ingredient := range lineIngredients {
			ingredients = append(ingredients, ingredient)
		}

		lineTimers, parsedLine, timerError := discoverTimers(parsedLine)
		if timerError != nil {
			return nil, timerError
		}
		for _, timer := range lineTimers {
			timers = append(timers, timer)
		}

		lineCookware, parsedLine := discoverCookware(parsedLine)
		for _, item := range lineCookware {
			cookware = append(cookware, item)
		}
		directions = append(directions, parsedLine)
	}
	recipe := &Recipe{
		Ingredients: ingredients,
		Directions:  directions,
		Timers:      timers,
		Cookware:    cookware,
	}
	return recipe, nil
}

func discoverIngredients(line string) ([]Ingredient, string, error) {
	var ingredients []Ingredient

	reQuantity := regexp.MustCompile(`\@([^\{\@#]+)\{((\d+\.?\d*)(\%[^\}]+)?)?\}`)
	for _, m := range reQuantity.FindAllStringSubmatch(line, -1) {
		name := m[1]

		var quantity float64
		if m[3] != "" {
			var err error
			quantity, err = strconv.ParseFloat(m[3], 64)
			if err != nil {
				return nil, "", err
			}
		} else {
			quantity = 1
		}

		unitPart := m[4]
		var unit string
		if len(unitPart) > 0 {
			unit = unitPart[1:]
		} else {
			unit = ""
		}

		ingredient := Ingredient{
			Name: name,
			Quantity: Quantity{
				Amount: QuantityAmount(quantity),
				Unit:   unit,
			},
		}
		ingredients = append(ingredients, ingredient)
	}

	line = reQuantity.ReplaceAllString(line, "$1")

	reNoQuantity := regexp.MustCompile(`\@(\pL+)`)
	for _, m := range reNoQuantity.FindAllString(line, -1) {
		name := m[1:]
		ingredient := Ingredient{
			Name: name,
			Quantity: Quantity{
				Amount: QuantityAmount(1),
			},
		}
		ingredients = append(ingredients, ingredient)
	}

	line = reNoQuantity.ReplaceAllString(line, "$1")

	return ingredients, line, nil
}

func removeComments(line string) string {
	commentRegex := regexp.MustCompile("//.*")
	commentsRemoved := commentRegex.ReplaceAllString(line, "")
	return strings.TrimSpace(commentsRemoved)
}

func discoverTimers(line string) ([]Timer, string, error) {
	var timers []Timer
	timerRegex := regexp.MustCompile(`~\{(\d+\.?\d*)%([^\}]+)\}`)
	for _, m := range timerRegex.FindAllStringSubmatch(line, -1) {
		duration, err := strconv.ParseFloat(m[1], 64)
		if err != nil {
			return nil, "", err
		}
		unit := m[2]
		timer := Timer{
			Duration: duration,
			Unit:     unit,
		}
		timers = append(timers, timer)
	}
	parsedLine := timerRegex.ReplaceAllString(line, "$1 $2")
	return timers, parsedLine, nil
}

func discoverCookware(line string) ([]string, string) {
	var cookware []string
	cookwareSpacesRegex := regexp.MustCompile(`#([^\{]+)\{\}`)
	for _, m := range cookwareSpacesRegex.FindAllStringSubmatch(line, -1) {
		name := m[1]
		cookware = append(cookware, name)
	}
	line = cookwareSpacesRegex.ReplaceAllString(line, "$1")

	cookwareNoSpacesRegex := regexp.MustCompile(`#(\pL+)`)
	for _, m := range cookwareNoSpacesRegex.FindAllStringSubmatch(line, -1) {
		name := m[1]
		cookware = append(cookware, name)
	}
	line = cookwareNoSpacesRegex.ReplaceAllString(line, "$1")

	return cookware, line
}
