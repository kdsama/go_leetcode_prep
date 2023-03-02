package grind75

import (
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	to_return := [][]int{}
	check := map[string]int{}
	for i := 0; i < len(nums)-1; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {

			if nums[left]+nums[right]+nums[i] == 0 {
				r := strconv.Itoa(nums[left]) + strconv.Itoa(nums[i]) + strconv.Itoa(nums[right])
				if _, ok := check[r]; ok {
					left++
					right--
					continue
				}
				to_return = append(to_return, []int{nums[left], nums[i], nums[right]})
				check[r] = 1
				left++
				right--
			} else if nums[left]+nums[right]+nums[i] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return to_return

}
