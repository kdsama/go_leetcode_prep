package grind75

// Given an integer n, return an array ans of length n + 1 such
//  that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary
// representation of i.

func CountBits(n int) []int {

	// 0 0 0
	// 1 1 1
	// 1 10 2
	// 2 11 3
	// 1 100 4
	// 2 101 5
	// 2 110 6
	// 3 111 7
	// 1 1000 8
	// 2 1001 9
	// 2 1010 10
	// 3 1011 11
	// 2 1100 12
	// 3 1101 13
	// 3 1110 14
	// 4 1111 15
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	if n == 2 {
		return []int{0, 1, 1}
	}
	arr := []int{0, 1, 1}

	for i := 3; i < n+1; i++ {
		k := getNumberOfOness(i)
		arr = append(arr, k)
	}
	return arr
}
func getNumberOfOness(n int) int {
	count := 0
	q := n
	r := 0

	for q != 0 {
		r = q % 2
		q = q / 2

		if r == 1 {
			count++
		}

	}
	return count
}
