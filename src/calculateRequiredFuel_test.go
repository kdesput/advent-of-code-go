package main

import "testing"

func TestRequiredFuelForMass(t *testing.T) {
	arguments := [4]int64{12, 14, 1969, 100756}
	expectedValues := [4]int64{2, 2, 654, 33583}

	for i := 0; i < 4; i++ {
		requiredFuel := requiredFuelForMass(arguments[i])
		if requiredFuel != expectedValues[i] {
			t.Errorf("calculateRequiredFuelForMass(%v) failed, expected %v, got %v", arguments[i], expectedValues[i], requiredFuel)
		} else {
			t.Logf("calculateRequiredFuelForMass(%v) passed", arguments[i])
		}
	}
}

func TestRequiredFuelForModule(t *testing.T) {
	arguments := [3]int64{14, 1969, 100756}
	expectedValues := [3]int64{2, 966, 50346}

	for i := 0; i < 3; i++ {
		requiredFuel := requiredFuelForModule(arguments[i])
		if requiredFuel != expectedValues[i] {
			t.Errorf("calculateRequiredFuelForModule(%v) failed, expected %v, got %v", arguments[i], expectedValues[i], requiredFuel)
		} else {
			t.Logf("calculateRequiredFuelForModule(%v) passed", arguments[i])
		}
	}
}
