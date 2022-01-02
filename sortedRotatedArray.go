package main

import (
	"fmt"
	"math"
)

func getMinimmumElement(slice []int64) int64 {
	// Returns the mimimum value from a sorted rotated array
	var (
		first int = 0
		last  int = len(slice) - 1
	)

	for true {
		var mid int = (first + last) / 2
		if first == last {
			return slice[first]
		} else if slice[mid] > slice[first] && slice[mid] > slice[last] {
			first = mid
		} else if slice[mid] < slice[first] && slice[mid] < slice[last] {
			last = mid
		} else {
			return int64(math.Min(float64(slice[first]), float64(slice[last])))
		}
	}
	return -1
}

func runGetMinimmumElement() {
	// Create a slice to hold elements
	var slice []int64 = []int64{6, 7, 1, 2, 3, 4, 5}

	// Print slice
	for _, val := range slice {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Find the minimum element
	fmt.Printf("Minimum element is %d\n", getMinimmumElement(slice))
}
