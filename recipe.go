package main

import "strings"

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
	var duration float64
	var unit string
	if t1.Unit == t2.Unit {
		unit = t1.Unit
		duration = t1.Duration + t2.Duration
	} else {
		d1 := durationToMinutes(t1)
		d2 := durationToMinutes(t2)
		duration = d1 + d2
		unit = "minutes"
	}
	return Timer{
		Duration: duration,
		Unit:     unit,
	}
}

func durationToMinutes(t *Timer) float64 {
	if strings.HasPrefix(t.Unit, "min") {
		return t.Duration
	}
	if strings.HasPrefix(t.Unit, "sec") {
		return t.Duration / 60
	}
	if strings.HasPrefix(t.Unit, "hour") {
		return t.Duration * 60
	}
	return 0
}
