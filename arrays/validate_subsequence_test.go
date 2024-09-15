package arrays

import (
	"testing"
)

func TestIsValidSubsequence(t *testing.T) {
	tests := []struct {
		array    []int
		sequence []int
		exResult bool
	}{
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 25, 6, -1, 8, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 6, -1, 8, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{22, 25, 6}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, -1, 8, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{25}, true},
		{[]int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 1, 1}, false},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 25, 6, -1, 8, 10, 12}, false},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{4, 5, 1, 22, 25, 6, -1, 8, 10}, false},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 23, 6, -1, 8, 10}, false},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 22, 25, 6, -1, 8, 10}, false},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 22, 6, -1, 8, 10}, false},
		{[]int{1, 2, 3, 4}, []int{2, 3}, true},                       // Simple case in the middle
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{}, true},           // Empty sequence
		{[]int{10, 20, 30, 40}, []int{20, 30, 40}, true},             // End of array subsequence
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{6, -1, 11}, false}, // Element not found
		{[]int{5, 5, 5, 5, 5}, []int{5, 5, 5}, true},                 // Repeated numbers
		{[]int{1, 2, 3, 4, 5}, []int{1, 3, 5}, true},                 // Non-consecutive subsequence
		{[]int{3, 1, 4, 2, 5}, []int{1, 4, 5}, true},                 // Non-sorted subsequence
		{[]int{1, 2, 3, 4}, []int{1, 2, 5}, false},                   // Subsequence longer than part of the array
		{[]int{7, 8, 9, 10}, []int{9, 10, 11}, false},                // Subsequence partially exists
		{[]int{0, 0, 0, 0}, []int{0, 0}, true},                       // Sequence with zeros
	}

	for i, test := range tests {
		got := IsValidSubsequence(test.array, test.sequence)

		if got != test.exResult {
			t.Errorf("Test case #%d failed: expected %v, got %v", i, test.exResult, got)
		}
	}
}
