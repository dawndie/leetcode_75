package main

import "math"

// quy về dạng bài tim min1, min2 từ 0 - n và kiểm tra xem min 1 và min 2 có bẻ hơn n hay ko.

func increasingTriplet(nums []int) bool {
	first := math.MaxInt
	second := math.MaxInt
	for _, n := range nums {
		if n <= first {
			first = n
		} else if n <= second {
			second = n
		} else {
			return true
		}
	}

	return false
}
