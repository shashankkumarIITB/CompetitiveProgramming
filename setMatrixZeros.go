// https://leetcode.com/problems/set-matrix-zeroes/

package main

import "fmt"

// Function to return a array of the given length initialized with zeros
func arrayOfZeros(length int) []int {
	arr := []int{}
	for i := 0; i < length; i++ {
		arr = append(arr, 0)
	}
	return arr
}

// Space complexity: O(1) for a (m x n) matrix
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	// Decide whether the corner is set to zero based on a row or column
	firstRow := false
	firstColumn := false
	// Loop over the matrix searching for zeros and populate the first row and first column with zeros
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Found zero at (i,j)
			if matrix[i][j] == 0 {
				if i == 0 {
					firstRow = true
				}
				if j == 0 {
					firstColumn = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// Set zeros in the rows
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			matrix[i] = arrayOfZeros(n)
		}
	}
	col := 1
	// Decide zeros for the first column
	if firstColumn == true {
		col = 0
	}
	// Set zeros in the columns
	for j := col; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	// Decide zeros for the first row
	if firstRow == true {
		matrix[0] = arrayOfZeros(n)
	}
}

func runSetMatrixZeros() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	fmt.Println(matrix)
	setZeroes(matrix)
	fmt.Println(matrix)
}
