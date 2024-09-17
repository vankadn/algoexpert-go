package arrays

import (
	"testing"
)

func TestNonConstructibleChange(t *testing.T) {
	tests := []struct {
		array    []int
		exResult int
	}{
		{[]int{1, 1, 1, 1}, 5},            // Multiple identical coins
		{[]int{1, 1, 3, 4}, 10},           // Mixed small and medium coins
		{[]int{5, 7, 1, 1, 2, 3, 22}, 20}, // Mixed small, medium, and large coins
		{[]int{1, 2, 3, 4, 9}, 20},        // Progressive set of coins
		{[]int{}, 1},                      // Empty array case

	}

	for i, test := range tests {
		//got := genAllSubsets(test.array, len(test.array))
		got := NonConstructibleChange(test.array)
		if got != test.exResult {
			t.Errorf("Test case #%d failed: expected %v, got %v", i, test.exResult, got)
		}
	}
}
