package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	mapTable := make(map[rune]int)
	longestSubstringLength := 0
	left, right := 0, 0
	for pos, char := range s {
		if value, ok := mapTable[char]; !ok || value < left {
			mapTable[char] = pos
		} else {

			left = mapTable[char] + 1
			mapTable[char] = pos
		}
		if count := (right - left) + 1; count > longestSubstringLength {
			longestSubstringLength = count
		}
		right++

	}
	return longestSubstringLength
}

// @lc code=end
