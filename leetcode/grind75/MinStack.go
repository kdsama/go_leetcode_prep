package grind75

type MinStack struct {
	arr []int
	min []int
}

// func Constructor() MinStack {
// 	return MinStack{[]int{}, []int{}}
// }

func (this *MinStack) Push(val int) {
	if len(this.min) == 0 {
		this.min = append(this.min, val)
	} else if val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
	this.arr = append(this.arr, val)
}

func (this *MinStack) Pop() {
	last := this.arr[len(this.arr)-1]

	if last == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}

	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
