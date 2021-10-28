package cuisiner

import (
	"regexp"
	"strconv"
	"strings"
)

func Parse(input string) (*Recipe, error) {
	var ingredients []Ingredient
	var directions []string
	var timers []Timer
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := removeComments(line)
		if len(trimmedLine) == 0 {
			// Skip empty lines
			continue
		}
		lineIngredients, parsedLine := discoverIngredients(trimmedLine)
		for _, ingredient := range lineIngredients {
			ingredients = append(ingredients, ingredient)
		}
		lineTimers, parsedLine := discoverTimers(parsedLine)
		for _, timer := range lineTimers {
			timers = append(timers, timer)
		}
		directions = append(directions, parsedLine)
	}
	recipe := &Recipe{
		Ingredients: ingredients,
		Directions:  directions,
		Timers:      timers,
	}
	return recipe, nil
}

func discoverIngredients(line string) ([]Ingredient, string) {
	var ingredients []Ingredient

	reQuantity := regexp.MustCompile(`\@([^\@]+)\{(\d+\.?\d*)(\%[^\}]+)?\}`)
	for _, m := range reQuantity.FindAllStringSubmatch(line, -1) {
		name := m[1]
		quantity, err := strconv.ParseFloat(m[2], 64)
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
		ingredient := Ingredient{
			Name:     name,
			Quantity: QuantityAmount(quantity),
			Unit:     unit,
		}
		ingredients = append(ingredients, ingredient)
	}

	line = reQuantity.ReplaceAllString(line, "$1")

	reNoQuantity := regexp.MustCompile(`\@(\w+)`)
	for _, m := range reNoQuantity.FindAllString(line, -1) {
		name := m[1:]
		ingredient := Ingredient{
			Name:     name,
			Quantity: QuantityAmount(1),
		}
		ingredients = append(ingredients, ingredient)
	}

	line = reNoQuantity.ReplaceAllString(line, "$1")

	return ingredients, line
}

func removeComments(line string) string {
	commentRegex := regexp.MustCompile("//.*")
	commentsRemoved := commentRegex.ReplaceAllString(line, "")
	return strings.TrimSpace(commentsRemoved)
}

func discoverTimers(line string) ([]Timer, string) {
	var timers []Timer
	timerRegex := regexp.MustCompile(`~\{(\d+\.?\d*)%([^\}]+)\}`)
	for _, m := range timerRegex.FindAllStringSubmatch(line, -1) {
		duration, err := strconv.ParseFloat(m[1], 64)
		if err != nil {
			panic("Bad duration") // TODO: Do not panic
		}
		unit := m[2]
		timer := Timer{
			Duration: duration,
			Unit:     unit,
		}
		timers = append(timers, timer)
	}
	parsedLine := timerRegex.ReplaceAllString(line, "$1 $2")
	return timers, parsedLine
}
