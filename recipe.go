package cuisiner

type QuantityAmount float64

type Ingredient struct {
	Name     string
	Quantity QuantityAmount
	Unit     string
}

type Recipe struct {
	Ingredients []Ingredient
	Directions  []string
	Timers      []Timer
}

type Timer struct {
	Duration float64
	Unit     string
}
