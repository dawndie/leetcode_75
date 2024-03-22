package week_1

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Value := 0
	l2Value := 0
	factor := 1
	for l1 != nil {
		l1Value += l1.Val * factor
		factor *= 10
		l1 = l1.Next
	}
	factor = 1
	for l2 != nil {
		l2Value += l2.Val * factor
		factor *= 10
		l2 = l2.Next
	}
	sumOfTwoNum := l1Value + l2Value
	splitedResult := splitInt(sumOfTwoNum)
	start := &ListNode{}
	current := start
	for i := 0; i < len(splitedResult); i++ {
		current.Next = &ListNode{}
		current = current.Next
		current.Val = splitedResult[i]
	}
	return start.Next
}

func splitInt(value int) []int {
	if value == 0 {
		return []int{0}
	}
	arr := []int{}
	for value > 0 {
		arr = append(arr, value%10)
		value /= 10
	}
	return arr
}

// @lc code=end
