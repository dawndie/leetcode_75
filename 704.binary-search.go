package main

/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	pivot := (low + high) / 2
	for low <= high {
		if nums[pivot] == target {
			return pivot
		}
		if  target < nums[pivot] {
			high = pivot - 1
		}
		if target > nums[pivot] {
			low = pivot + 1
		}
		pivot = (low + high) / 2
	}
	return -1

}

// @lc code=end
