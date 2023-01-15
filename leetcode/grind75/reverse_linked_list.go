package grind75

import (
	"fmt"

	"algorithms.com/m/leetcode/data_structures"
)

func ReverseList(head *data_structures.ListNode) *data_structures.ListNode {
	tmp := data_structures.ListNode{}
	first := head
	for head != nil {
		tmp = *head.Next
		head.Next = nil
		tmp.Next = head
		fmt.Println(head.Val)
		head = head.Next

	}
	return first
}
