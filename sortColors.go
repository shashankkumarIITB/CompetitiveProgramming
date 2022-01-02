// https://leetcode.com/problems/sort-colors/

package main

import "fmt"

// Bubble sort
func sortColors(nums []int) {
	swapped := true
	for {
		if swapped {
			swapped = false
			for i := 0; i < len(nums)-1; i++ {
				if nums[i] > nums[i+1] {
					swapped = true
					nums[i], nums[i+1] = nums[i+1], nums[i]
				}
			}
		} else {
			break
		}
	}
}

func runSortColors() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
