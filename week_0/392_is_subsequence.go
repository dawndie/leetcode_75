package main

// s is sequence of t
func isSubsequence(s string, t string) bool {
	sequenceIndex := 0
	sequenceLen := len(s)
	if s == "" {
		return true
	} else if t == "" {
		return false
	}

	for i := 0; i < len(t); i++ {
		if t[i] == s[sequenceIndex] {
			sequenceIndex++
		}
		if sequenceLen == sequenceIndex {
			return true
		}

	}
	return false
}
