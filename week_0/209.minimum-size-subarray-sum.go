package main

import (
	"math"
)

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	left, right, sum := 0, 0, 0
	min := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		right++
		for sum >= target && left < right {
			if count := right - left; count < min {
				min = count
			}
			sum -= nums[left]
			left++
		}
		if left == right && sum >= target {
			return 1
		}
	}
	if left == 0 && right == len(nums) && sum < target {
		return 0
	}
	return min
}

// @lc code=end
