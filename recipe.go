package cuisiner

type Ingredient struct {
	Name     string
	Quantity int
	Unit string
}

type Recipe struct {
	Ingredients []Ingredient
	Directions  []string
}
