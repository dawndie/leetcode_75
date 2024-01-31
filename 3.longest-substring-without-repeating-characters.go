package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	max := 0
	mapTable := make(map[rune]int)
	for pos, char := range s {
		if value, ok := mapTable[char]; !ok || value < left {
			mapTable[char] = pos
		} else {
			mapTable[char] = pos
			left = value + 1
		}
		right++
		if right-left > max {
			max = right - left
		}
	}
	return max
}

// @lc code=end
