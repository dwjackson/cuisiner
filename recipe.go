package main

type Recipe struct {
	Ingredients []Ingredient
	Directions  []string
	Timers      []Timer
	Cookware    []string
}

func (r *Recipe) IngredientsList() []Ingredient {
	return r.Ingredients
}

type Timer struct {
	Duration float64
	Unit     string
}

func (t1 *Timer) Add(t2 *Timer) Timer {
	duration := t1.Duration + t2.Duration
	return Timer{
		Duration: duration,
		Unit:     "minutes",
	}
}
