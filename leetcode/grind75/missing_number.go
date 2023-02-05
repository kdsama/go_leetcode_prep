package grind75

import "fmt"

func MissingNumber(nums []int) int {
	// sort and check
	// nlogn + n

	sum := 0
	len_nums := len(nums)
	// 3 means 0 , 1 , 2
	for i := range nums {
		sum += nums[i]
	}
	expected := ((len_nums) * (len_nums + 1)) / 2
	// we are expecting the sum to be 3
	fmt.Println(expected, sum)
	// but the sum is 4

	if expected != sum {
		if expected > sum {
			return expected - sum
		} else {
			return sum - expected
		}
	} else {
		if sum == (len_nums-2*len_nums-1)/2 {
			return len(nums)
		}
	}
	return 0

}
