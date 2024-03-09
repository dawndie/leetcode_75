package week_0

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	output := [][]int{}
	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			output = append(output, newInterval)
			output = append(output, intervals[i:]...)
			return output
		} else if newInterval[0] > intervals[i][1] {
			output = append(output, intervals[i])
		} else {
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
		}
	}
	output = append(output, newInterval)
	return output
}

// @lc code=end
