package main

import "reflect"

/*
 * @lc app=leetcode id=1657 lang=golang
 *
 * [1657] Determine if Two Strings Are Close
 */

// @lc code=start
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	charMap1 := make(map[rune]int)
	charMap2 := make(map[rune]int)
	occurrencesMap1 := make(map[int]int)
	occurrencesMap2 := make(map[int]int)
	for _, char := range word1 {
		charMap1[char]++
	}
	for _, char := range word2 {
		charMap2[char]++
	}
	if len(charMap1) != len(charMap2) {
		return false
	}
	for key, value := range charMap1 {
		if _, ok := charMap2[key]; !ok {
			return false
		}
		occurrencesMap1[value]++
	}
	for _, value := range charMap2 {
		occurrencesMap2[value]++
	}
	equal := reflect.DeepEqual(occurrencesMap1, occurrencesMap2)
	if equal {
		return true
	}
	return false
}

// @lc code=end
