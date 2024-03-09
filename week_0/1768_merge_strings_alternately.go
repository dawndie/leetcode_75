package main

func mergeAlternately(word1 string, word2 string) string {
	length := 0
	mergedStr := ""
	lengthOfW1 := len(word1)
	lengthOfW2 := len(word2)
	if lengthOfW1 > lengthOfW2 {
		length = lengthOfW1
	} else {
		length = lengthOfW2
	}
	for i := 0; i < length; i++ {
		if i < lengthOfW1 {
			mergedStr += string(word1[i])
		}

		if i < lengthOfW2 {
			mergedStr += string(word2[i])
		}
	}
	return mergedStr
}
