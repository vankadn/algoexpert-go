package arrays

import "sort"

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	runningCount := 0
	sort.Ints(coins)

	for _, coin := range coins {
		if runningCount+1 < coin {
			return runningCount + 1
		}
		runningCount = runningCount + coin
	}
	return runningCount + 1
}
