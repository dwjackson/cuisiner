package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Timer struct {
	duration float64
	Unit     string
}

func (t *Timer) DurationString() string {
	isInteger := math.Floor(t.duration) == t.duration
	var s string
	if isInteger {
		s = strconv.Itoa(int(t.duration))
	} else {
		s = fmt.Sprintf("%5.5f", t.duration)
	}
	s = strings.TrimRight(s, "0")
	return s
}

func (t1 *Timer) Add(t2 *Timer) Timer {
	var duration float64
	var unit string
	if t1.Unit == t2.Unit {
		unit = t1.Unit
		duration = t1.duration + t2.duration
	} else {
		d1 := durationToMinutes(t1)
		d2 := durationToMinutes(t2)
		duration = d1 + d2
		unit = "minutes"
	}
	return Timer{
		duration: duration,
		Unit:     unit,
	}
}

func durationToMinutes(t *Timer) float64 {
	if strings.HasPrefix(t.Unit, "min") {
		return t.duration
	}
	if strings.HasPrefix(t.Unit, "sec") {
		return t.duration / 60
	}
	if strings.HasPrefix(t.Unit, "hour") {
		return t.duration * 60
	}
	return 0
}
