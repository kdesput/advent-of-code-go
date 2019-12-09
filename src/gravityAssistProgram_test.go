package main

import "testing"

func TestGravityAssistProgram(t *testing.T) {
	arguments := [4][]int{[]int{1, 0, 0, 0, 99}, []int{2, 3, 0, 3, 99}, []int{2, 4, 4, 5, 99, 0}, []int{1, 1, 1, 4, 99, 5, 6, 0, 99}}
	expectedValues := [4]int{2, 2, 2, 30}

	for i := 0; i < 4; i++ {
		valueAtPosition0 := gravityAssistProgram(arguments[i])
		if valueAtPosition0 != expectedValues[i] {
			t.Errorf("gravityAssistProgram(%v) failed, expected %v, got %v", arguments[i], expectedValues[i], valueAtPosition0)
		} else {
			t.Logf("gravityAssistProgram(%v) passed", arguments[i])
		}
	}
}
