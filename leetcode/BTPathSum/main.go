package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil{
		return false
	}
	if targetSum == root.Val && root.Left == nil && root.Right == nil{
		return true
	}
	targetSum = targetSum - root.Val
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 7}
	// root.Right = &TreeNode{Val: 3}
	// root.Right.Left = &TreeNode{Val: 5}
	// root.Right.Right = &TreeNode{Val: 6}

	fmt.Println(hasPathSum(&root, 100))
}
