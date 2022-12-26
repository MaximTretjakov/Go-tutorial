package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convert(nums []int, low, height int) *TreeNode {
	if low > height {
		return nil
	}
	mid := low + (height-low)/2
	node := new(TreeNode)
	node.Val = nums[mid]
	node.Left = convert(nums, low, mid-1)
	node.Right = convert(nums, mid+1, height)
	return node
}

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	return convert(nums, 0, len(nums)-1)
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(nums).Val)
}
