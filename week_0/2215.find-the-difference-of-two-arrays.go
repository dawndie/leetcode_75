package main

/*
 * @lc app=leetcode id=2215 lang=golang
 *
 * [2215] Find the Difference of Two Arrays
 */

// @lc code=start
func findDifference(nums1 []int, nums2 []int) [][]int {
	mapTable := make(map[int]int)
	result := make([][]int, 2)
	for i := 0; i < len(nums1); i++ {
		if _, ok := mapTable[nums1[i]]; !ok {
			mapTable[nums1[i]] = 1
		}

	}
	for i := 0; i < len(nums2); i++ {
		if value, ok := mapTable[nums2[i]]; !ok {
			mapTable[nums2[i]] = 2
		} else if value == 1 {
			mapTable[nums2[i]] = 3
		}
	}
	for key, value := range mapTable {
		if value == 1 {
			result[0] = append(result[0], key)
		}
		if value == 2 {
			result[1] = append(result[1], key)
		}
	}
	return result
}

// @lc code=end
