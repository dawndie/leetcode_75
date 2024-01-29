package main

/*
 * @lc app=leetcode id=724 lang=golang
 *
 * [724] Find Pivot Index
 */

// @lc code=start
func pivotIndex(nums []int) int {
	length := len(nums)
	sumOfLeft := make([]int, length)
	sumOfRight := make([]int, length)
	currentSumOfLeft, currentSumOfRight := 0, 0
	for i := 0; i < length; i++ {
		sumOfLeft[i] = currentSumOfLeft
		currentSumOfLeft += nums[i]
		sumOfRight[length-i-1] = currentSumOfRight
		currentSumOfRight += nums[length-i-1]
	}
	for i := 0; i < length; i++ {
		if sumOfLeft[i] == sumOfRight[i] {
			return i
		}
	}
	return -1
}

// @lc code=end
