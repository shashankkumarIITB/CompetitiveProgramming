// https://leetcode.com/problems/merge-sorted-array/

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for {
		if j < 0 {
			// Traversed nums2 => exit
			break
		} else if i < 0 && j >= 0 {
			// Traversed nums1 but not nums2 => move nums2 to nums1
			for ; j > -1; j-- {
				nums1[j] = nums2[j]
			}
		} else if nums1[i] > nums2[j] {
			// Move the largest number to end of nums1
			nums1[i+j+1], i = nums1[i], i-1
		} else {
			// Move the largest number to end of nums1
			nums1[i+j+1], j = nums2[j], j-1
		}
	}
}

func runMerge() {
	nums1, nums2 := []int{4, 5, 6, 0, 0, 0}, []int{1, 2, 3}
	m, n := 3, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
