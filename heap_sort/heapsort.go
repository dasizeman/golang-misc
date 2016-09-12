package main

import (
	"fmt"
	"github.com/dasizeman/tools"
)

func heapsort(A []int) {
	buildMaxHeap(A)
	heapSize := len(A)

	for i := len(A); i > 1; i-- {
		swap(A, 1, i)
		heapSize--
		maxHeapify(A, 1, heapSize)
	}
}

func buildMaxHeap(A []int) {
	for i := len(A) / 2; i >= 1; i-- {
		maxHeapify(A, i, len(A))
	}
}

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func maxHeapify(A []int, i, heapSize int) {
	l := left(i)
	r := right(i)

	largest := i
	if l <= heapSize && A[l-1] > A[largest-1] {
		largest = l
	}
	if r <= heapSize && A[r-1] > A[largest-1] {
		largest = r
	}

	if largest != i {
		swap(A, i, largest)
		maxHeapify(A, largest, heapSize)
	}
}

func swap(A []int, i, j int) {
	i--
	j--
	temp := A[i]
	A[i] = A[j]
	A[j] = temp
}

func main() {
	test := tools.RandomIntSlice(1, 50, 50)
	fmt.Printf("Input: %v\n", test)
	heapsort(test)
	fmt.Printf("Result: %v\n", test)
}
