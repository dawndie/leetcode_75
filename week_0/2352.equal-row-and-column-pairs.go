package main

/*
 * @lc app=leetcode id=2352 lang=golang
 *
 * [2352] Equal Row and Column Pairs
 */

// @lc code=start
func equalPairs(grid [][]int) int {
	length := len(grid)
	m := make(map[[200]int]int)
	arr := [200]int{}
	for i := 0; i < length; i++ {
		copy(arr[:], grid[i])
		m[arr]++
	}
	res := 0
	for i := 0; i < length; i++ {
		arr = [200]int{}
		for j := 0; j < length; j++ {
			arr[j] = grid[j][i]
		}
		if value, ok := m[arr]; ok {
			res += value
		}
	}
	return res
}

// @lc code=end
