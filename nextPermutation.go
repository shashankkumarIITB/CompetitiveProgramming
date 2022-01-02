// https://leetcode.com/problems/next-permutation/
package main

import (
	"fmt"
)

// Returns the next lexicographic permutation of the sequence
// Time complexity = O(n)
// Space complexity = O(1)
func nextPermutation(nums []int) {
	switchedNums := false
	// Last index after which the sequence is decreasing
	lastDecreasingIndex := len(nums) - 1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			// Revert all the numbers after this as they are in the decreasing order
			for j, k := lastDecreasingIndex, len(nums)-1; j < k; j, k = j+1, k-1 {
				nums[j], nums[k] = nums[k], nums[j]
			}
			// Replace the i-th number with the number just larger than it
			for j := i; j < len(nums); j++ {
				if nums[i-1] < nums[j] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					break
				}
			}
			switchedNums = true
			break
		} else {
			lastDecreasingIndex -= 1
		}
	}
	// The sequence already represents the largest number
	// Reverse the sequence to get the smallest sequence
	if !switchedNums {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

func runNextPermutation() {
	nums := []int{2, 1, 4, 3}
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
}
