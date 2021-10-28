package main

import (
	"fmt"
	"math"
	"io/ioutil"
	"os"
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
	contentBytes, readError := ioutil.ReadFile(fileName)
	if readError != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	content := string(contentBytes)
	recipe, parseError := Parse(content)
	if parseError != nil {
		fmt.Println("Error parsing recipe")
		os.Exit(1)
	}

	fmt.Println("Ingredients")
	for _, ingredient := range recipe.Ingredients {
		q := float64(ingredient.Quantity)
		if math.Floor(q) == q && q > 1.0 {
			qInt := int(q)
			fmt.Printf("\t* %d%s %s\n", qInt, ingredient.Unit, ingredient.Name)
		} else if q > 1.0 {
			fmt.Printf("\t* %.2f%s %s\n", ingredient.Quantity, ingredient.Unit, ingredient.Name)
		} else {
			fmt.Printf("\t* %s\n", ingredient.Name)
		}
	}
	// TODO
}
