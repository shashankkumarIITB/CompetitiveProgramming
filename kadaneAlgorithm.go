// https://leetcode.com/problems/maximum-subarray/

package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	maxSum, currentSum := math.MinInt64, 0
	for _, val := range nums {
		currentSum += val
		if currentSum > maxSum {
			maxSum = currentSum
		}
		if currentSum < 0 {
			currentSum = 0
		}
	}
	return maxSum
}

func runMaxSubArray() {
	fmt.Println(maxSubArray([]int{0, -3, -2, -3, -2, 2, -3, 0, 1, -1}))
	fmt.Println(maxSubArray([]int{-3, -2, -3, -2, -3, -1}))
}
