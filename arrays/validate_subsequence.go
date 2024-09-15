package arrays

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	i := 0
	j := 0
	for i < len(array) && j < len(sequence) {
		fmt.Printf("seq value %d array value %d \n", sequence[j], array[i])
		if array[i] == sequence[j] {
			j++
		}
		i++
	}
	return j == len(sequence)
}
