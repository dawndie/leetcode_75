package main

/*
 * @lc app=leetcode id=1732 lang=golang
 *
 * [1732] Find the Highest Altitude
 */

// @lc code=start
func largestAltitude(gain []int) int {
	highestAltitude, currentAltitude := 0, 0
	for i := 0; i < len(gain); i++ {
		currentAltitude = currentAltitude + gain[i]
		if highestAltitude < currentAltitude {
			highestAltitude = currentAltitude
		}
	}
	return highestAltitude
}

// @lc code=end
