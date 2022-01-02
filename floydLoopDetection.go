// https://leetcode.com/problems/find-the-duplicate-number/

package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow, fast = nums[slow], nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	for {
		slow = 0
		slow, fast = nums[slow], nums[fast]
		if slow == fast {
			return slow
		}
	}
}

func main() {
	fmt.Println(findDuplicate([]int{4, 1, 3, 4, 2}))
}
