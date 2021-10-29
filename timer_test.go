package main

import "testing"

func TestAddTimersWithSameUnit(t *testing.T) {
	t1 := Timer{
		Duration: 30,
		Unit:     "minutes",
	}

	t2 := Timer{
		Duration: 15,
		Unit:     "minutes",
	}

	result := t1.Add(&t2)
	expected := Timer{
		Duration: 45,
		Unit:     "minutes",
	}
	testTimer(t, &expected, &result)
}

func TestAddTimersWithDifferentUnits(t *testing.T) {
	t1 := Timer{
		Duration: 30,
		Unit:     "minutes",
	}

	t2 := Timer{
		Duration: 1,
		Unit:     "hour",
	}

	result := t1.Add(&t2)
	expected := Timer{
		Duration: 90,
		Unit:     "minutes",
	}
	testTimer(t, &expected, &result)
}

func testTimer(t *testing.T, expected *Timer, actual *Timer) {
	if actual.Duration != expected.Duration {
		t.Fatalf("Expected %f, got %f", expected.Duration, actual.Duration)
	}
	if actual.Unit != expected.Unit {
		t.Fatalf("Expected unit to be %s, was %s", expected.Unit, actual.Unit)
	}
}
