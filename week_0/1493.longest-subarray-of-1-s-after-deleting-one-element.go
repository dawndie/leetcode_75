package main

/*
 * @lc app=leetcode id=1493 lang=golang
 *
 * [1493] Longest Subarray of 1's After Deleting One Element
 */

// @lc code=start
func longestSubarray(nums []int) int {
	length := len(nums)
	l, r := 0, 0
	count := 0
	res := 0
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			count++
		}
		for count > 1 {
			if nums[l] == 0 {
				count--
			}
			l++
		}
		if res < r-l {
			res = r - l
		}
		r++
	}
	return res
}

// @lc code=end
