// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package main

import "fmt"

func maxProfit(prices []int) int {
	nums := make([]int, len(prices)-1)
	for i := 0; i < len(prices)-1; i++ {
		nums = append(nums, prices[i+1]-prices[i])
	}
	// Use Kadane's algorithm to find maximum sum of subarray
	return maxSubArray(nums)
}

func runMaxProfit() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
