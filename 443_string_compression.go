package main

import "fmt"

func compress(chars []byte) int {
	writeIndex := 0
	count := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			count++
		} else {
			chars[writeIndex] = chars[i-1]
			writeIndex++
			if count > 1 {
				for _, digit := range []byte(fmt.Sprintf("%d", count)) {
					chars[writeIndex] = digit
					writeIndex++
				}
			}
			count = 1
		}
	}

	chars[writeIndex] = chars[len(chars)-1]
	writeIndex++
	if count > 1 {
		for _, digit := range []byte(fmt.Sprintf("%d", count)) {
			chars[writeIndex] = digit
			writeIndex++
		}
	}

	return writeIndex
}
