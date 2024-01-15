package main

func moveZeroes(nums []int) {
	movedIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[movedIndex] = nums[i]
			nums[i] = 0
			movedIndex++
		}
	}
	for i := movedIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}
