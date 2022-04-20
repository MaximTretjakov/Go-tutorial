package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			l2 = nil
			continue
		}
		if l2 == nil {
			cur.Next = l1
			l1 = nil
			continue
		}

		if l1.Val > l2.Val {
			cur.Next = l2
			cur, l2 = cur.Next, l2.Next
		} else {
			cur.Next = l1
			cur, l1 = cur.Next, l1.Next
		}
	}
	return res.Next
}

func main() {
	// t1 := "t1"
	// t2 := "t2"
	// q1 := "q1"
	// q2 := "q2"

	// fmt.Printf("Before t1: %s t2: %s\n", t1, t2)
	// fmt.Printf("Before q1: %s q2: %s\n", q1, q2)

	// t1, t2 = q1, q2

	// fmt.Printf("After t1: %s t2: %s\n", t1, t2)
	// fmt.Printf("After q1: %s q2: %s\n", q1, q2)
	ln2 := &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 100, Next: nil}}}
	cur1 := &ListNode{}
	// fmt.Printf("Empty cur1: Val: %T\t Next: %T\n", cur1.Val, cur1.Next)
	cur1.Next = ln2
	// fmt.Printf("cur1.Next = ln2: %T\t\n", cur1.Next)
	cur1, ln2 = cur1.Next, ln2.Next
	fmt.Printf("cur1 = cur1.Next: Val: %d\t Next: %T\n", cur1.Val, cur1.Next)
	fmt.Printf("ln2 = ln2.Next: Val: %d\t Next: %d\n", ln2.Val, ln2.Next.Val)
}
