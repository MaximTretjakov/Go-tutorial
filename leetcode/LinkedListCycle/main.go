package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil{
        return false
    }
    slow := head
    fast := head.Next
    
    for fast.Next != nil && fast.Next.Next != nil && fast != slow{
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    return slow == fast
}

func main() {
	// cycled
	l := ListNode{Val: 1}
	cycle := &l
	l.Next = &ListNode{Val: 2}
	l.Next.Next = &ListNode{Val: 3, Next: cycle}
	// non cycled
	// s := ListNode{Val: 1}
	// s.Next = &ListNode{Val: 3}
	// s.Next.Next = &ListNode{Val: 111, Next: nil}
	// // one node
	// t:=ListNode{Val: 333, Next: nil}

	fmt.Println(hasCycle(&l))
	// fmt.Println(hasCycle(&s))
	// fmt.Println(hasCycle(&t))
}
