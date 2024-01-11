package main

import (
	"strings"
	"unicode"
)

func printString(str string) string {
	duplicateTime := 0
	result := ""
	for _, char := range str {
		if unicode.IsDigit(char) {
			parsedIntChar := int(char - '0')
			duplicateTime = duplicateTime*10 + parsedIntChar
		} else {
			repeatedString := strings.Repeat(string(char), duplicateTime)
			result += repeatedString
			duplicateTime = 0
		}
	}
	return result
}
