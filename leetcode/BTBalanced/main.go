package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    var dfs func(*TreeNode, int) (int, bool)
    dfs = func(n *TreeNode, height int) (int, bool) {
        if n == nil {
            return height, true
        }
        leftHeight, leftBalanced := dfs(n.Left, height+1)
        rightHeight, rightBalanced := dfs(n.Right, height+1)
        return max(leftHeight, rightHeight), leftBalanced && rightBalanced && abs(leftHeight - rightHeight) <= 1
    }

    _, ans := dfs(root, 0)
    return ans    
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
	// positive
	pos := TreeNode{Val: 1}
	pos.Right = &TreeNode{Val: 2}
	pos.Right.Left = &TreeNode{Val: 44}
	pos.Right.Right = &TreeNode{Val: 88}
	pos.Left = &TreeNode{Val: 9}
	// negative
	neg := TreeNode{Val: 1}
	neg.Right = &TreeNode{Val: 2}
	neg.Left = &TreeNode{Val: 2}
	neg.Left.Left = &TreeNode{Val: 3}
	neg.Left.Right = &TreeNode{Val: 3}
	neg.Left.Left.Left = &TreeNode{Val: 4}
	neg.Left.Left.Right = &TreeNode{Val: 4}

	fmt.Println(isBalanced(&pos))
}
