/*
3. Fast and Slow Pointers
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			break
		}
		slow = slow.Next
	}
	return slow
}

func main() {
	s := ListNode{Val: 1}
	s.Next = &ListNode{Val: 3}
	s.Next.Next = &ListNode{Val: 111}
	s.Next.Next.Next = &ListNode{Val: 222}
	s.Next.Next.Next.Next = &ListNode{Val: 333, Next: nil}
	res := middleNode(&s)
	fmt.Printf("Address: %v Val: %d Next: %v\n", res, res.Val, res.Next)
}
