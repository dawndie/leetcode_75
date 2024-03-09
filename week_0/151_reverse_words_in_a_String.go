package week_0

func reverseWords(s string) string {
	word := ""
	wordList := []string{}
	for pos, char := range s {
		if char != ' ' {
			word += string(char)
		}
		if (char == ' ' || pos == len(s)-1) && word != "" {
			wordList = append(wordList, word)
			word = ""
		}
	}
	result := ""
	for i := len(wordList) - 1; i >= 0; i-- {
		result = result + wordList[i]
		if i > 0 {
			result += " "
		}
	}
	return result
}
