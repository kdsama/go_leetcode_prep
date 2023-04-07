package graph

// if we find all the nodes that dont have no incoming edges, it means other edges have an incoming edges.

func FindSmallestSetOfVertices(n int, edges [][]int) []int {
	// first get adjaceny matrix from the edges
	adjList := map[int][]int{}
	incomingList := map[int][]int{}

	// visited means probably dont need to go any further.
	to_return := []int{}
	for i := range edges {
		adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
		incomingList[edges[i][1]] = append(incomingList[edges[i][0]], edges[i][0])

	}
	// the below loops atleast gives us the nodes with no incoming edges. We need to include them
	for i := 0; i < n; i++ {
		if _, ok := incomingList[i]; !ok {
			to_return = append(to_return, i)
		}
	}

	// first find the nodes with no incoming edge
	//	adjMatrix = node --> outgoing edges
	// we need to find node that are not present as incoming edges
	// means that are not present on the right side at all .

	return to_return
}
