package grind75

func ProductExceptSelf(nums []int) []int {
	// without division, and O(n)
	arr1 := make([]int, len(nums))
	arr2 := make([]int, len(nums))
	p1 := 1
	p2 := 1
	n := len(nums) - 1

	to_return := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		p1 = p1 * nums[i]
		p2 = p2 * nums[n-i]
		arr1[i] = p1
		arr2[n-i] = p2
	}
	for i := 1; i < n; i++ {
		to_return[i] = arr1[i-1] * arr2[i+1]
	}
	to_return[0] = arr2[1]
	to_return[n] = arr1[n-1]
	return to_return
}
