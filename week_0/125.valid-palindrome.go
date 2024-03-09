package main

import (
	"regexp"
	"strings"
)

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	pattern := "[^a-z0-9]"
	s = regexp.MustCompile(pattern).ReplaceAllString(s, "")
	i, j := 0, len(s)-1
	if len(s) < 2 {
		return true
	}
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end
