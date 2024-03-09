package week_0

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for position, element := range nums {
		if _, ok := hashMap[element]; ok {
			return []int{hashMap[element], position}
		}
		complement := target - element
		hashMap[complement] = position
	}
	return []int{0, 0}
}

// @lc code=end
