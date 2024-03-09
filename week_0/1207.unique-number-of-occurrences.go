package week_0

/*
 * @lc app=leetcode id=1207 lang=golang
 *
 * [1207] Unique Number of Occurrences
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	intMap := make(map[int]int)
	occurrencesMap := make(map[int]int)
	for _, value := range arr {
		intMap[value]++
	}
	for key, value := range intMap {
		if _, ok := occurrencesMap[value]; ok {
			return false
		}
		occurrencesMap[value] = key
	}
	return true
}

// @lc code=end
