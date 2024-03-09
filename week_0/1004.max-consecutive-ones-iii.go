package week_0

/*
 * @lc app=leetcode id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	length := len(nums)
	l, r := 0, 0
	count := 0
	res := 0
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			count++
		}

		for count > k {
			if nums[l] == 0 {
				count--
			}
			l++
		}
		r++
		if r-l > res {
			res = r - l
		}
	}
	return res
}

// @lc code=end
