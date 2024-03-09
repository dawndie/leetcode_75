package week_0

func maxOperations(nums []int, k int) int {
	numsMap := make(map[int]int)
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= k {
			continue
		}
		remaining := k - nums[i]
		if value, ok := numsMap[nums[i]]; ok && value > 0 {
			count++
			numsMap[nums[i]]--
		} else {
			numsMap[remaining]++
		}
	}
	return count
}
