package week_0

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := maxHeight(root.Left)
	rightHeight := maxHeight(root.Right)
	if leftHeight-rightHeight > 1 || leftHeight-rightHeight < -1 {
		return false
	}
	return isBalanced(root.Left) == true && isBalanced(root.Right) == true
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := maxHeight(root.Left) + 1
	rightHeight := maxHeight(root.Right) + 1
	if leftHeight > rightHeight {
		return leftHeight
	} else {
		return rightHeight
	}
}

// @lc code=end
