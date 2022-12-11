package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	isSame := false

	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val == q.Val {
			isSame = isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}
	
	return isSame
}

func main() {
	p := TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 3}

	q := TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 2}
	q.Right = &TreeNode{Val: 3}

	fmt.Printf("is same tree?\t%v\n", isSameTree(&p, &q))
}
