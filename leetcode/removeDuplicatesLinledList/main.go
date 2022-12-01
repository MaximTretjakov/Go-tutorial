package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := head

	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}
	return dummyHead
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{},
				},
			},
		},
	}

	l := deleteDuplicates(list)
	fmt.Println(l.Val)
	fmt.Println(l.Next.Val)
	fmt.Println(l.Next.Next.Val)
}
