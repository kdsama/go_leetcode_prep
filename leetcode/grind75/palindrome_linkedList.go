package grind75

import (
	"algorithms.com/m/leetcode/data_structures"
)

func isPalindrome(head *data_structures.ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	l := len(arr) - 1
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[l-i] {
			return false
		}
	}
	return true
}
