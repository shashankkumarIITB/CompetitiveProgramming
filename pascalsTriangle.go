// https://leetcode.com/problems/pascals-triangle/

package main

import "fmt"

// Function to generate Pascal's Triangle
func generate(numRows int) [][]int {
	pascalTriangle := [][]int{}
	// Base case
	if numRows == 1 {
		return [][]int{{1}}
	} else {
		// Start with numRows = 2
		pascalTriangle = [][]int{{1}, {1, 1}}
		for i := 2; i < numRows; i++ {
			arr := []int{1}
			// Build iteratively using the last row in the triangle
			lastRow := pascalTriangle[len(pascalTriangle)-1]
			for j := 0; j < len(lastRow)-1; j++ {
				arr = append(arr, lastRow[j]+lastRow[j+1])
			}
			arr = append(arr, 1)
			// Append the new row to the triangle
			pascalTriangle = append(pascalTriangle, arr)
		}
	}
	return pascalTriangle
}

func runGenerate() {
	fmt.Println(generate(5))
}
