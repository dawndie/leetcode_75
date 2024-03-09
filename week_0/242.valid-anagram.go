package week_0

import "reflect"

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}
	hashMapS := make(map[byte]int)
	hashMapT := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hashMapS[s[i]]++
		hashMapT[t[i]]++
	}
	return reflect.DeepEqual(hashMapS, hashMapT)
}
// @lc code=end

