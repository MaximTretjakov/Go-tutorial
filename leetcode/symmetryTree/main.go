package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(left, right *TreeNode) bool {
	switch {
	case left == nil && right == nil:
		return true
	case left == nil && right != nil, left != nil && right == nil:
		return false
	default:
		if left.Val != right.Val {
			return false
		} else {
			return bfs(left.Left, right.Right) && bfs(left.Right, right.Left)
		}
	}
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return bfs(root.Left, root.Right)
}
