package main

func reverseVowels(s string) string {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	i, j := 0, len(s)-1
	bytes := []byte(s)
	for i < j {
		if !vowels[bytes[i]] {
			i++
			continue
		}
		if !vowels[bytes[j]] {
			j--
			continue
		}
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
