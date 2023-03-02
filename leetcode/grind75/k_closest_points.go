package grind75

import "container/heap"

type Point struct {
	x        int
	y        int
	distance int
}
type MyHeap []Point

func (h MyHeap) Len() int           { return len(h) }
func (h MyHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h MyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MyHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}
func (h *MyHeap) Pop() interface{} {
	// Just returning the last one and removing it as well
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func KClosest(points [][]int, k int) [][]int {
	h := &MyHeap{}
	heap.Init(h)
	returning_arr := [][]int{}
	for i := range points {
		heap.Push(h, Point{points[i][0], points[i][1], points[i][0]*points[i][0] + points[i][1]*points[i][1]})
	}

	for i := 0; i < k; i++ {
		smallestPointInHeap := heap.Pop(h).(Point)
		returning_arr = append(returning_arr, []int{smallestPointInHeap.x, smallestPointInHeap.y})
	}
	return returning_arr
}
