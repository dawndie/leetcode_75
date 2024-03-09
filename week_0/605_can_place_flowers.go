package week_0

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	count := 0
	i := 0

	for i < length {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == length-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
			i += 2
		} else {
			i++
		}
	}

	return count >= n
}
