package main

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	stack := []string{}
	// subString := []string{}
	// for i := 0; i < len(s); i++ {
	// 	if s[i] != ']' {
	// 		stack = append(stack, string(s[i]))
	// 	} else {
	// 		subString = nil
	// 		for stack[len(stack)] != "[" {
	// 			subString = append(subString, stack[len(stack)])
	// 			stack = stack[:len(stack)-1]
	// 		}
	// 		multiplizer := stack[len(stack)-2]
	// 		stack = stack[:len(stack)-2]
			

	// 	}
	// }
	return stack[0]
}

// @lc code=end
