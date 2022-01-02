package main

import (
	"fmt"
	"sort"
)

func containsNTimes(nums []int, num int, freq int) bool {
	// Returns true if num is present in nums for freq times
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == num {
			freq -= 1
		} else {
			break
		}
		if freq == 0 {
			return true
		}
	}
	return false
}

func subsetsWithDup(nums []int) [][]int {
	// Array to store results
	result := [][]int{[]int{}}
	if len(nums) > 0 {
		// Given that numbers are in the range [-10, 10],
		// holds value for the last seen number
		last_num := 11
		last_num_freq := 0
		// Sort the numbers
		sort.Ints(nums)
		// Loop over the numbers
		for _, val := range nums {
			// Array to hold the next sets to be added to the result
			var arr [][]int
			if val == last_num {
				// Append only to the sets where last number == val
				for i := len(result) - 1; i >= 0; i-- {
					if containsNTimes(result[i], val, last_num_freq) {
						temp := append(result[i], val)
						arr = append(arr, temp)
					} else {
						break
					}
				}
				// Update last number's frequency
				last_num_freq += 1
			} else {
				// Append to all the previous sets
				for i := 0; i < len(result); i++ {
					temp := append(result[i], val)
					arr = append(arr, temp)
				}
				// Update the last seen number and its frequency
				last_num = val
				last_num_freq = 1
			}
			// Append subsets from this iteration to result
			fmt.Println(result)
			result = append(result, arr...)
			fmt.Println(result)
		}
	}
	return result
}

func runSubsetsWithDup() {
	nums := []int{1, 1, 2, 2, 3}
	result := subsetsWithDup(nums)
	fmt.Println(result)
}
