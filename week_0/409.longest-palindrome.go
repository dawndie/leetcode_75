package week_0

/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome(s string) int {
	count := 0
	isOneCharExist := false
	alphabetCountTable := make(map[string]int)
	for _, v := range s {
		alphabetCountTable[string(v)]++
	}
	for k, v := range alphabetCountTable {
		if v/2 > 0 {
			count = count + (v/2)*2
			alphabetCountTable[k] = v % 2
		}
		if alphabetCountTable[k] == 1 {
			isOneCharExist = true
		}
	}
	if isOneCharExist {
		return count + 1
	}
	return count
}

// @lc code=end
