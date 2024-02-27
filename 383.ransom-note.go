package main

/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	alphabetCountTable := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		alphabetCountTable[magazine[i]-'a']++
	}
	for _, v := range ransomNote {
		if alphabetCountTable[v-'a'] > 0 {
			alphabetCountTable[v-'a']--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
