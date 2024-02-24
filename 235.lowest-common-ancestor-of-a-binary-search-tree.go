package main

/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else {
			return cur
		}
	}
	return nil
}

// @lc code=end
