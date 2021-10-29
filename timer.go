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

func (t *Timer) ToString() string {
	var hours, minutes, seconds float64
	minutes = durationToMinutes(t)
	if minutes >= 60 {
		hours = math.Floor(minutes / 60)
		minutes = minutes - hours*60
	}
	if !isInteger(minutes) {
		totalSeconds := minutes * 60
		minutes = math.Floor(totalSeconds / 60)
		seconds = totalSeconds - minutes*60
	}

	var parts []string
	if hours > 0 {
		var hoursUnit string
		if hours == 1 {
			hoursUnit = "hour"
		} else {
			hoursUnit = "hours"
		}
		hoursString := formatDuration(hours) + " " + hoursUnit
		parts = append(parts, hoursString)
	}
	if minutes > 0 {
		var minutesUnit string
		if minutes == 1 {
			minutesUnit = "minute"
		} else {
			minutesUnit = "minutes"
		}
		minutesString := formatDuration(minutes) + " " + minutesUnit
		parts = append(parts, minutesString)
	}
	if seconds > 0 {
		var secondsUnit string
		if seconds == 1 {
			secondsUnit = "second"
		} else {
			secondsUnit = "seconds"
		}
		secondsString := formatDuration(seconds) + " " + secondsUnit
		parts = append(parts, secondsString)
	}
	durationString := strings.Join(parts, ", ")
	return durationString
}

func formatDuration(duration float64) string {
	var durationString string
	if isInteger(duration) {
		durationString = strconv.Itoa(int(duration))
	} else {
		durationString = fmt.Sprintf("%5.5f", duration)
		durationString = strings.TrimRight(durationString, "0")
	}
	return durationString
}

func isInteger(f float64) bool {
	return math.Floor(f) == f
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
	if isMinutes(t.Unit) {
		return t.duration
	}
	if isSeconds(t.Unit) {
		return t.duration / 60
	}
	if isHours(t.Unit) {
		return t.duration * 60
	}
	return 0
}

func isMinutes(unit string) bool {
	return strings.HasPrefix(unit, "min")
}

func isSeconds(unit string) bool {
	return strings.HasPrefix(unit, "sec")
}

func isHours(unit string) bool {
	return strings.HasPrefix(unit, "hour")
}
