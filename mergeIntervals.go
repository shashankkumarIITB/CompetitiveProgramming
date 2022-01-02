// https://leetcode.com/problems/merge-intervals/

package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// Sort the intervals with respect to the starting point
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// Loop over and merge the overlapping intervals
	mergedIntervals := make([][]int, 0, len(intervals))
	interval := []int{-1, -1}
	for _, val := range intervals {
		if val[0] > interval[1] {
			// Add the interval to the merged interval list
			if interval[0] != -1 {
				mergedIntervals = append(mergedIntervals, interval)
			}
			// Start a new interval
			interval = val
		} else {
			// Modify the end of the interval
			if val[1] > interval[1] {
				interval[1] = val[1]
			}
		}
	}
	// Append the last interval
	return append(mergedIntervals, interval)
}

func runMerge() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
