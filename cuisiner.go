package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const USAGE string = "USAGE: cuisiner [COMMAND] [ARGS...]"

type Command = func([]string)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(USAGE)
		os.Exit(1)
	}
	commandName := os.Args[1]
	commands := map[string]Command{
		"print":    printCommand,
		"shopping": shoppingCommand,
	}
	if command, commandExists := commands[commandName]; commandExists {
		command(os.Args[2:])
	} else {
		fmt.Printf("Invalid command: %s\n", commandName)
		os.Exit(1)
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
	for _, ingredient := range recipe.IngredientsList() {
		ingredientLine := createIngredientLine(&ingredient)
		fmt.Printf("* %s\n", ingredientLine)
	}
	fmt.Println("")

	if len(recipe.Cookware) > 0 {
		fmt.Println("## Cookware")
		fmt.Println("")
		for _, item := range recipe.Cookware {
			fmt.Printf("* %s\n", item)
		}
		fmt.Println("")
	}

	if len(recipe.Timers) > 0 {
		fmt.Println("## Total Time")
		total := recipe.Timers[0]
		for i, timer := range recipe.Timers {
			if i == 0 {
				// Skip the first
				continue
			}
			total = total.Add(&timer)
		}
		isInteger := math.Floor(total.Duration) == total.Duration
		var duration string
		if isInteger {
			duration = strconv.Itoa(int(total.Duration))
		} else {
			duration = fmt.Sprintf("%f", total.Duration)
			duration = strings.Trim(duration, "0")
			if duration[0] == '.' {
				duration = "0" + duration
			}
		}
		fmt.Printf("\n%s %s\n\n", duration, total.Unit)
	}

	fmt.Println("## Directions")
	fmt.Println("")
	for i, direction := range recipe.Directions {
		fmt.Printf("%d. %s\n", i+1, direction)
	}
}

func parseRecipeFile(fileName string) (*Recipe, error) {
	content, readError := readFileToString(fileName)
	if readError != nil {
		return nil, readError
	}
	recipe, parseError := Parse(content)
	if parseError != nil {
		return nil, parseError
	}
	return recipe, nil
}

func readFileToString(fileName string) (string, error) {
	contentBytes, readError := ioutil.ReadFile(fileName)
	if readError != nil {
		return "", errors.New("Error reading file")
	}
	content := string(contentBytes)
	return content, nil
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

func shoppingCommand(args []string) {
	var pantry *Pantry
	if len(args) > 0 {
		pantryFileName := args[0]
		var pantryError error
		pantry, pantryError = parsePantryFile(pantryFileName)
		if pantryError != nil {
			fmt.Printf("%s\n", pantryError)
			os.Exit(1)
		}
	}

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

	list := ShoppingList(recipes, pantry)

	fmt.Println("# Shopping List")
	fmt.Println("")
	for _, item := range list {
		fmt.Printf("* %s\n", item)
	}
}

func createIngredientLine(ingredient *Ingredient) string {
	amount := float64(ingredient.Quantity.Amount)
	isInteger := math.Floor(amount) == amount
	name := ingredient.Name
	unit := ingredient.Quantity.Unit

	if isInteger {
		amountInt := int(amount)
		if unit != "" {
			return fmt.Sprintf("%d %s %s", amountInt, unit, name)
		}
		if amount != 1 {
			return fmt.Sprintf("%d %s", amountInt, name)
		}
		return fmt.Sprintf("%s", name)
	}

	amountStr := fmt.Sprintf("%.5f", ingredient.Quantity.Amount)
	amountStr = strings.Trim(amountStr, "0")
	if amountStr[0] == '.' {
		amountStr = "0" + amountStr
	}

	if unit != "" {
		return fmt.Sprintf("%s %s %s", amountStr, unit, name)
	}

	return fmt.Sprintf("%s %s", amountStr, name)
}

func parsePantryFile(fileName string) (*Pantry, error) {
	pantryContent, readError := readFileToString(fileName)
	if readError != nil {
		return nil, readError
	}
	var ingredients []Ingredient
	for _, line := range strings.Split(pantryContent, "\n") {
		lineIngredients, _, err := parseIngredients(line)
		if err != nil || len(lineIngredients) == 0 {
			continue
		}
		for _, ingredient := range lineIngredients {
			ingredients = append(ingredients, ingredient)
		}
	}
	pantry := &Pantry{
		ingredients,
	}
	return pantry, nil
}
