package leetcode

// You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
// The digits are ordered from most significant to least significant in left-to-right order.
//  The large integer does not contain any leading 0's.

// Increment the large integer by one and return the resulting array of digits.

func PlusOne(digits []int) []int {

	len_d := len(digits)
	// s := -1
	carryOver := -1
	for i := len_d - 1; i >= 0; i-- {
		s := digits[i]
		s += 1
		if s >= 10 {
			carryOver = s / 10
			s = s % 10
			digits[i] = s
		} else {
			digits[i] = s
			carryOver = -1
			break
		}

	}
	if carryOver != -1 {
		digits = append([]int{carryOver}, digits...)
	}
	return digits

}
