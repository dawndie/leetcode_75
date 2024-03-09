package main

/*
 * @lc app=leetcode id=997 lang=golang
 *
 * [997] Find the Town Judge
 */

// @lc code=start
func findJudge(n int, trust [][]int) int {
	l, r := make([]int, n+1), make([]int, n+1)
	for _, t := range trust {
		l[t[0]]++
		r[t[1]]++
	}
	for i := 1; i <= n; i++ {
		if l[i] == 0 && r[i] == n-1 {
			return i
		}
	}
	return -1
}

// @lc code=end
