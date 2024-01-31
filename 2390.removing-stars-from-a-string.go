package main

/*
 * @lc app=leetcode id=2390 lang=golang
 *
 * [2390] Removing Stars From a String
 */

// @lc code=start
type Stack []rune

func (s *Stack) Push(n rune) {
	*s = append(*s, n)
}

func (s *Stack) Pop() int {
	length := len(*s)
	if length == 0 {
		return -1
	}
	n := length - 1
	*s = (*s)[:n]
	return n
}

func removeStars(s string) string {
	stack := Stack{}
	for _, char := range s {
		if char == '*' {
			stack.Pop()
		} else {
			stack.Push(char)
		}
	}
	return string(stack)
}

// @lc code=end
