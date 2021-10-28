package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

const USAGE string = "USAGE: cuisiner [COMMAND] [ARGS...]"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(USAGE)
		os.Exit(1)
	}
	commandName := os.Args[1]
	switch commandName {
	case "print":
		printCommand(os.Args[2:])
	case "html":
		htmlCommand(os.Args[2:])
	default:
		fmt.Printf("Invalid command: %s\n", commandName)
	}
}

func printCommand(args []string) {
	if len(args) < 1 {
		fmt.Println("File name required")
		os.Exit(1)
	}
	fileName := args[0]
	recipe, err := parseRecipeFile(fileName)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Println("Ingredients")
	for _, ingredient := range recipe.Ingredients {
		q := float64(ingredient.Quantity.Amount)
		if math.Floor(q) == q && q > 1.0 {
			qInt := int(q)
			fmt.Printf("\t* %d%s %s\n", qInt, ingredient.Quantity.Unit, ingredient.Name)
		} else if q > 1.0 {
			fmt.Printf("\t* %.2f%s %s\n", ingredient.Quantity.Amount, ingredient.Quantity.Unit, ingredient.Name)
		} else {
			fmt.Printf("\t* %s\n", ingredient.Name)
		}
	}

	fmt.Println("")
	fmt.Println("Directions")
	for i, direction := range recipe.Directions {
		fmt.Printf("\t%d. %s\n", i+1, direction)
	}
}

func parseRecipeFile(fileName string) (*Recipe, error) {
	contentBytes, readError := ioutil.ReadFile(fileName)
	if readError != nil {
		return nil, errors.New("Error reading file")
	}
	content := string(contentBytes)
	recipe, parseError := Parse(content)
	if parseError != nil {
		return nil, errors.New("Error parsing recipe")
	}
	return recipe, nil
}

func htmlCommand(args []string) {
	if len(args) < 1 {
		fmt.Println("File name required")
		os.Exit(1)
	}

	fileName := args[0]
	recipe, err := parseRecipeFile(fileName)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	lastDotIndex := strings.LastIndex(fileName, ".")
	recipeName := fileName[0:lastDotIndex]

	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println("  <head>")
	fmt.Println("    <meta charset=\"utf-8\">")
	fmt.Printf("    <title>%s</title>\n", recipeName)
	fmt.Println("  </head>")
	fmt.Println("  <body>")
	fmt.Println("    <h1>Ingredients</h1>")
	fmt.Println("    <ul>")
	for _, ingredient := range recipe.Ingredients {
		q := float64(ingredient.Quantity.Amount)
		if math.Floor(q) == q && q > 1.0 {
			qInt := int(q)
			fmt.Printf("      <li>%d%s %s</li>\n", qInt, ingredient.Quantity.Unit, ingredient.Name)
		} else if q > 1.0 {
			fmt.Printf("      <li>%.2f%s %s</li>\n", ingredient.Quantity.Amount, ingredient.Quantity.Unit, ingredient.Name)
		} else {
			fmt.Printf("<li>%s</li>\n", ingredient.Name)
		}
	}
	fmt.Println("    </ul>")
	fmt.Println("    <h1>Directions</h1>")
	fmt.Println("    <ol>")
	for _, direction := range recipe.Directions {
		fmt.Printf("      <li>%s</li>\n", direction)
	}
	fmt.Println("    </ol>")
	fmt.Println("  </body>")
	fmt.Println("</html>")
}
