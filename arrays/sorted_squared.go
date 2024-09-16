package arrays

import "fmt"

// https://go.dev/play/p/_Y6GiQ8vocv
func SortedSquaredArray(array []int) []int {
	// Write your code here.
	left := 0
	right := len(array) - 1
	index := len(array) - 1
	newArray := make([]int, len(array))
	for left <= right {
		x := array[left] * array[left]
		y := array[right] * array[right]
		fmt.Printf("comparing sqr of i %d x %d and  sqr of j %d y %d \n", left, array[left], right, array[right])
		if x > y {
			fmt.Printf("adding i  %d sqr to arr \n", left)
			newArray[index] = x
			left++
		} else {
			fmt.Printf("adding j  %d sqr to arr \n", right)
			newArray[index] = y
			right--
		}
		fmt.Printf("new output array %v \n", newArray)
		index--
	}
	return newArray
}
