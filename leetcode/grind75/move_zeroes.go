package grind75

func MoveZeroes(nums []int) {

	counter := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[counter] = nums[i]
			counter++

		}
	}

	for k := counter; k < len(nums); k++ {
		nums[k] = 0
	}

}
