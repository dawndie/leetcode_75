package week_0

/*
 * @lc app=leetcode id=2390 lang=golang
 *
 * [2390] Removing Stars From a String
 */

// @lc code=start
type RuneStack []rune

func (s *RuneStack) Push(n rune) {
	*s = append(*s, n)
}

func (s *RuneStack) Pop() int {
	length := len(*s)
	if length == 0 {
		return -1
	}
	n := length - 1
	*s = (*s)[:n]
	return n
}

func removeStars(s string) string {
	stack := RuneStack{}
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
