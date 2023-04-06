package graph

import "fmt"

var final [][]int

func AllPathsSourceTarget(graph [][]int) [][]int {
	// bfs
	// the graph is a queue for bfs.
	// need to convert it into paths.
	arr := []int{}

	AllPathHelper(graph, 0, arr)
	return final
}

func AllPathHelper(graph [][]int, el int, arr []int) {

	arr = append(arr, el)
	if el >= len(graph)-1 {
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println(arr)
		final = append(final, arr)
		fmt.Println(final[len(final)-1])
		return
	}
	for _, v := range graph[el] {

		AllPathHelper(graph, v, arr)

	}
}
