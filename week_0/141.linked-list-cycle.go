package week_0

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	linkedListMap := make(map[*ListNode]bool)
	pointer := head
	for pointer != nil {
		if linkedListMap[pointer] == false {
			linkedListMap[pointer] = true
			pointer = pointer.Next
		} else {
			return true
		}
	}
	return false
}

// @lc code=end
