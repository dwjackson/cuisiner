package cuisiner

import "strings"

type Recipe struct {
	Directions []string
}

func Parse(input string) (*Recipe, error) {
	var directions []string
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			directions = append(directions, line)
		}
	}
	recipe := &Recipe{
		Directions: directions,
	}
	return recipe, nil
}
