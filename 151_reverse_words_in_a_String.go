package main

import "fmt"

func reverseWords(s string) string {
	word := ""
	wordList := []string{}
	for pos, char := range s {
		if char != ' ' || pos == len(s)-1 {
			word += string(char)
		} else {
			wordList = append(wordList, word)
			fmt.Print(wordList)
			word = ""
		}
	}
	result := ""
	for i := len(wordList) - 1; i > 0; i-- {
		result = result + wordList[i] + " "

	}
	return result
}
