package main

import (
	"bufio"
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
	case "shopping":
		shoppingCommand()
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

	recipeTitle := recipeTitleFromFileName(fileName)

	recipe, err := parseRecipeFile(fileName)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("# %s\n\n", recipeTitle)

	fmt.Println("## Ingredients")
	fmt.Println("")
	for _, ingredient := range recipe.Ingredients {
		ingredientLine := createIngredientLine(&ingredient)
		fmt.Printf("* %s\n", ingredientLine)
	}
	fmt.Println("")

	fmt.Println("## Directions")
	fmt.Println("")
	for i, direction := range recipe.Directions {
		fmt.Printf("%d. %s\n", i+1, direction)
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

func recipeTitleFromFileName(fileName string) string {
	lastDotIndex := strings.LastIndex(fileName, ".")
	if lastDotIndex < 0 {
		lastDotIndex = len(fileName)
	}
	recipeTitle := fileName[0:lastDotIndex]
	recipeTitle = strings.ReplaceAll(recipeTitle, "_", " ")
	return recipeTitle
}

func shoppingCommand() {
	var recipes []Recipe
	reader := bufio.NewReader(os.Stdin)
	fileName, err := reader.ReadString('\n')
	for err == nil {
		fileName = strings.TrimSpace(fileName)
		if len(fileName) > 0 {
			recipe, err := parseRecipeFile(fileName)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}
			recipes = append(recipes, *recipe)
		}
		fileName, err = reader.ReadString('\n')
	}

	list := ShoppingList(recipes)

	fmt.Println("# Shopping List")
	fmt.Println("")
	for _, item := range list {
		fmt.Printf("* %s\n", item)
	}
}

func createIngredientLine(ingredient *Ingredient) string {
	q := float64(ingredient.Quantity.Amount)
	isInteger := math.Floor(q) == q

	if isInteger && ingredient.Quantity.Unit != "" {
		qInt := int(q)
		return fmt.Sprintf("%d %s %s", qInt, ingredient.Quantity.Unit, ingredient.Name)
	}

	if isInteger && q > 1 {
		qInt := int(q)
		return fmt.Sprintf("%d %s", qInt, ingredient.Name)
	}

	amountStr := fmt.Sprintf("%.5f", ingredient.Quantity.Amount)
	amountStr = strings.Trim(amountStr, "0")

	if q > 1 && ingredient.Quantity.Unit != "" {
		return fmt.Sprintf("%s %s %s", amountStr, ingredient.Quantity.Unit, ingredient.Name)
	}

	if q > 1 {
		return fmt.Sprintf("%s %s", amountStr, ingredient.Name)
	}

	return fmt.Sprintf("%s", ingredient.Name)
}
