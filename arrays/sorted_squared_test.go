package arrays

import (
	"testing"
)

func TestSortedSquaredArray(t *testing.T) {
	tests := []struct {
		array    []int
		exResult []int
	}{
		{[]int{0}, []int{0}},                                                         // Single zero element
		{[]int{1}, []int{1}},                                                         // Single positive element
		{[]int{-1}, []int{1}},                                                        // Single negative element, square becomes positive
		{[]int{-2, -1, 0, 1, 2}, []int{0, 1, 1, 4, 4}},                               // Mixed negative, positive, and zero
		{[]int{-5, -4, -3, -2, -1}, []int{1, 4, 9, 16, 25}},                          // All negative numbers
		{[]int{1, 2, 3, 4, 5}, []int{1, 4, 9, 16, 25}},                               // All positive numbers
		{[]int{-10, -3, 1, 5, 9}, []int{1, 9, 25, 81, 100}},                          // Mixed negative and positive
		{[]int{-10, -5, 0, 3, 8}, []int{0, 9, 25, 64, 100}},                          // Mixed negative and positive with zero
		{[]int{-7, -3, 0, 3, 7}, []int{0, 9, 9, 49, 49}},                             // Symmetrical around zero
		{[]int{-3, -3, -2, 2, 3, 3}, []int{4, 4, 9, 9, 9, 9}},                        // Array with duplicates
		{[]int{-1000, -100, 0, 100, 1000}, []int{0, 10000, 10000, 1000000, 1000000}}, // Very large numbers
		{[]int{-1, -1, -1, -1, -1}, []int{1, 1, 1, 1, 1}},                            // All elements the same
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},                                 // Array with all zeros
		{[]int{-3, 2, 4}, []int{4, 9, 16}},                                           // Small array with mixed signs
		{[]int{-9, -7, -5, 1, 4, 6}, []int{1, 16, 25, 36, 49, 81}},                   // Another mixed negative and positive
	}

	for i, test := range tests {
		got := SortedSquaredArray(test.array)

		if !Equal(got, test.exResult) {
			t.Errorf("Test case #%d failed: expected %v, got %v", i, test.exResult, got)
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
