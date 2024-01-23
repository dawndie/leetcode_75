package main

func maxOperations(nums []int, k int) int {
	numsMap := make(map[int]int)
	count := 0
	for i := 0; i < len(nums); i++ {
		if value, oke := numsMap[nums[i]]; oke && value != 0 {
			numsMap[nums[i]]--
			count++
			continue
		}
		if _, ok := numsMap[k-nums[i]]; ok {
			numsMap[k-nums[i]]++
		} else {
			numsMap[k-nums[i]] = 0
		}
	}
	return count
}
