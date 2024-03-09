package week_0

/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root.Right
	root.Right = root.Left
	root.Left = temp
	invertTree(root.Left)
	invertTree(root.Right)
}

// @lc code=end
