package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		if area := (min(height[left], height[right]) * (right - left)); area > maxArea {
			maxArea = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}
