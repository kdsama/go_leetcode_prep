package graph

import "fmt"

// n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
func ValidPath(n int, edges [][]int, source int, destination int) bool {
	mapping := map[int][]int{}
	traversed := map[int]bool{}
	for i := range edges {
		mapping[edges[i][0]] = append(mapping[edges[i][0]], edges[i][1])
		mapping[edges[i][1]] = append(mapping[edges[i][1]], edges[i][0])
	}
	fmt.Println(mapping)
	res := false
	for key, _ := range mapping {
		if !traversed[key] {
			res = res || valid_path_dfs(mapping, source, destination, traversed)
		}
	}
	return res
}

func valid_path_dfs(mapping map[int][]int, source int, destination int, traversed map[int]bool) bool {
	if source == destination {
		return true
	}
	traversed[source] = true
	to_return := false
	for i := range mapping[source] {

		if destination == mapping[source][i] {
			return true
		} else {
			if !traversed[mapping[source][i]] {
				to_return = to_return || valid_path_dfs(mapping, mapping[source][i], destination, traversed)
			}
		}

	}

	return to_return

}
