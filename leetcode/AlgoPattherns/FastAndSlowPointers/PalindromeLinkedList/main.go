/*
3. Fast and Slow Pointers
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast, stack := head, head, make([]int, 0)
	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil && fast.Next == nil {
		slow = slow.Next
	}
	for slow != nil {
		temp := stack[len(stack)-1]
		if slow.Val != temp {
			return false
		}
		stack = stack[:len(stack)-1]
		slow = slow.Next
	}
	return true
}

func main() {
	s := ListNode{Val: 1}
	s.Next = &ListNode{Val: 2}
	s.Next.Next = &ListNode{Val: 2}
	s.Next.Next.Next = &ListNode{Val: 1}
	fmt.Println(isPalindrome(&s))
}
