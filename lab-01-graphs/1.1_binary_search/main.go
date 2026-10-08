package main

import (
	"fmt"
	"slices"
)

func binarySearch(slice []int, target int) int {
	l := 0
	r := len(slice) - 1
	count := 0

	for l <= r {
		count++
		mid := l + (r-l)/2
		currNum := slice[mid]

		switch {
		case currNum == target:
			return count
		case currNum < target:
			l = mid + 1
		default:
			r = mid - 1
		}
	}

	return -1
}

func main() {
	testSlice := []int{9, 1, 5, 3, 9, 6}
	target := 6

	slices.Sort(testSlice)

	steps := binarySearch(testSlice, target)
	fmt.Printf("Количество шагов, чтобы найти число %d равно %d\n", target, steps)
}
