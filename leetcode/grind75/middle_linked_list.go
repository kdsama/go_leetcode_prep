package grind75

import "algorithms.com/m/leetcode/data_structures"

func middleNode(head *data_structures.ListNode) *data_structures.ListNode {

	tmp := head
	counter := 0
	pos := 0
	for head != nil {
		head = head.Next
		counter++
		for pos < counter/2 {
			tmp = tmp.Next
			pos++
		}
	}
	return tmp
}

// Alternate method
/*

slow := head
for head.Next != nil || head != nil {
	head = head.Next.Next
	slow = slow.Next
}
return slow

*/
