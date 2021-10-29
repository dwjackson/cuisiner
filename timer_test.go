package main

import "testing"

func TestAddTimersWithSameUnit(t *testing.T) {
	t1 := Timer{
		duration: 30,
		Unit:     "minutes",
	}

	t2 := Timer{
		duration: 15,
		Unit:     "minutes",
	}

	result := t1.Add(&t2)
	expected := Timer{
		duration: 45,
		Unit:     "minutes",
	}
	testTimer(t, &expected, &result)
}

func TestAddTimersWithDifferentUnits(t *testing.T) {
	t1 := Timer{
		duration: 30,
		Unit:     "minutes",
	}

	t2 := Timer{
		duration: 1,
		Unit:     "hour",
	}

	result := t1.Add(&t2)
	expected := Timer{
		duration: 90,
		Unit:     "minutes",
	}
	testTimer(t, &expected, &result)
}

func testTimer(t *testing.T, expected *Timer, actual *Timer) {
	if actual.duration != expected.duration {
		t.Fatalf("Expected %f, dot %f", expected.duration, actual.duration)
	}
	if actual.Unit != expected.Unit {
		t.Fatalf("Expected unit to be %s, was %s", expected.Unit, actual.Unit)
	}
}
