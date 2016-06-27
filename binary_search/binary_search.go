package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var amt, max int
	fmt.Printf("How many values?\n")
	_, err := fmt.Scanf("%d\n", &amt)
	fmt.Printf("Max value?\n")
	_, err = fmt.Scanf("%d\n", &max)

	values := randomSlice(amt, max)

	fmt.Printf("Generated: %v\n", values)

	fmt.Printf("Enter something to search for:\n")
	var search int
	_, err = fmt.Scanf("%d\n", &search)
	if err != nil {
		fmt.Printf("Something wrong with that input.\n")
	}

	res, idx := binarySearch(values, search)

	if res {
		fmt.Printf("%d was found at index %d.\n", search, idx)
	} else {
		fmt.Printf("%d was not found in the input.\n", search)
	}

}

func randomSlice(len, max int) (res []int) {
	for i := 0; i < len; i++ {
		res = append(res, rand.Intn(max))
	}
	return res
}

func binarySearch(a []int, search int) (bool, int) {
	fmt.Printf("\n\nStarting search...\n")
	// Make sure the array is sorted
	sort.Ints(a)
	fmt.Printf("Sorted input: %v\n", a)
	leftBound := -1
	rightBound := len(a)

	for leftBound+1 < rightBound {
		inspectIdx := leftBound + ((rightBound - leftBound) / 2)

		if search == a[inspectIdx] {
			return true, inspectIdx
		} else if search < a[inspectIdx] {
			rightBound = inspectIdx
		} else {
			leftBound = inspectIdx
		}
	}
	fmt.Printf("Finished search.\n")

	return false, -1
}
