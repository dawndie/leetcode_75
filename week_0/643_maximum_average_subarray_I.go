package main

func findMaxAverage(nums []int, k int) float64 {
	sum := float64(0)
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	avg := sum / float64(k)
	max := avg
	if k == len(nums) {
		return max
	}
	for i := 0; i < len(nums)-k; i++ {
		avg = ((avg * float64(k)) - float64(nums[i]) + float64(nums[i+k])) / float64(k)
		if avg > max {
			max = avg
		}
	}
	return max
}
