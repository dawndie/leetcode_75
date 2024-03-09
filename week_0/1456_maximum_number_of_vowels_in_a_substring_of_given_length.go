package main

func maxVowels(s string, k int) int {
	vowelsMap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	currentNumOfVowels := 0
	for i := 0; i < k; i++ {
		if _, ok := vowelsMap[s[i]]; ok {
			currentNumOfVowels++
		}
	}
	maxVowels := currentNumOfVowels
	if k == len(s) {
		return maxVowels
	}
	for i := 0; i < len(s)-k; i++ {
		if _, ok := vowelsMap[s[i]]; ok {
			currentNumOfVowels--
		}
		if _, ok := vowelsMap[s[i+k]]; ok {
			currentNumOfVowels++
		}
		if maxVowels < currentNumOfVowels {
			maxVowels = currentNumOfVowels
		}
	}
	return maxVowels
}
