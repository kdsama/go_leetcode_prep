package grind75

import (
	"algorithms.com/m/leetcode/linkedlist"
)

// 21. Merge Two Sorted Lists
// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.

func MergeTwoLists() {

	list1 := linkedlist.InsertAndReturnFirst([]int{5, 6, 7})
	list2 := linkedlist.InsertAndReturnFirst([]int{2, 3, 4})

	first := &linkedlist.ListNode{}
	list3 := first
	for list1 != nil && list2 != nil {
		if list2.Val <= list1.Val {

			list3.Next = list2

			list2 = list2.Next
			list3 = list3.Next
		} else {

			list3.Next = list1
			list1 = list1.Next
			list3 = list3.Next
		}

	}

	// Just attach the remainging linkedlist items at the end of new linked list l3 .
	// so if list1 items are remaining, we l3.next = first all remaining l1 linked list tems
	if list1 != nil {
		list3.Next = list1
	} else if list2 != nil {
		list3.Next = list2
	}
	linkedlist.PrintFromFirst(first.Next)
}
