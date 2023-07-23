package grind75

import "sort"

func merge(intervals [][]int) [][]int {
	new_arr := [][]int{}
	sort.Slice(intervals, func(a, b int) bool {
		return (intervals[a][0] < intervals[b][0]) || ((intervals[a][0] == intervals[b][0]) && (intervals[a][1] < intervals[b][1]))
	})
	i := 0
	for i < len(intervals) {
		curr := intervals[i]

		start := curr[0]
		end := curr[1]
		for i+1 < len(intervals) {
			curr = intervals[i+1]

			if end < curr[0] {

				break
			}
			if !(end > curr[1]) {
				end = curr[1]
			}
			i++
		}
		new_arr = append(new_arr, []int{start, end})
		i++
	}
	return new_arr
}
