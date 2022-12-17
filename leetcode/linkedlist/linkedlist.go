package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertAndReturnFirst(li []int) *ListNode {
	first := &ListNode{}
	it := first
	for i := range li {
		it.Next = &ListNode{Val: li[i], Next: nil}
		it = it.Next
	}
	return first.Next
}

func PrintFromFirst(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
