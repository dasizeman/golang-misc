package main

import (
	"fmt"
	"github.com/dasizeman/tools"
)

func merge(A, B []int) []int {
	merged := make([]int, len(A)+len(B))

	// Index into subarrays
	i := 0
	j := 0

	// Populate the merged array
	for k := 0; k < len(merged); k++ {

		// Handle cases where subarrays are used up
		if i > len(A)-1 { // Left array is "empty"
			merged[k] = B[j]
			j++
			continue
		}

		if j > len(B)-1 { // Right array is "empty"
			merged[k] = A[i]
			i++
			continue
		}

		if A[i] <= B[j] {
			merged[k] = A[i]
			i++
		} else {
			merged[k] = B[j]
			j++
		}
	}

	return merged
}

func mergesort(A []int) []int {
	if len(A) == 1 {
		return A
	}

	midIdx := len(A) / 2

	sortedLeft := mergesort(A[0:midIdx])
	sortedRight := mergesort(A[midIdx:len(A)])

	return merge(sortedLeft, sortedRight)
}

func main() {
	// Generate a test slice
	testSlice := tools.RandomIntSlice(1, 10, 10)
	sorted := mergesort(testSlice)

	fmt.Printf("Input:%v\nSorted:%v\n", testSlice, sorted)

}
