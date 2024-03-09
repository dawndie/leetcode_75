package week_0

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		candies[i] += extraCandies
		if candies[i] >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}
