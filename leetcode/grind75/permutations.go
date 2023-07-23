package grind75

func permute(nums []int) [][]int {

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	perms := [][]int{}
	for i, n := range nums {
		// create a len -1 arr
		pi := []int{}
		arr := make([]int, len(nums)-1)
		// copy content
		copy(arr[:i], nums[:i])
		copy(arr[i:], nums[i+1:])

		subperms := permute(arr)
		for _, p := range subperms {
			pi = append([]int{n}, p...)
			perms = append(perms, pi)
		}

	}
	return perms

}
