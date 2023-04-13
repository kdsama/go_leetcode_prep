package graph

import "fmt"

func FindOrder(numCourses int, prerequisites [][]int) []int {
	to_return := []int{}
	visited := map[int]bool{}
	adj := map[int][]int{}
	added := map[int]bool{}
	queue := []int{}

	for i := range prerequisites {
		adj[prerequisites[i][0]] = append(adj[prerequisites[i][0]], prerequisites[i][1])
	}
	fmt.Println(adj)
	if len(prerequisites) == 0 {
		for i := 0; i < numCourses; i++ {
			to_return = append(to_return, i)
		}
		return to_return
	}
	for i := 0; i < numCourses; i++ {
		queue = append(queue, i)

		if !findorder_bfs(visited, added, adj, &queue, &to_return) {
			return []int{}
		}

	}

	return to_return
}

func findorder_bfs(visited map[int]bool, added map[int]bool, adj map[int][]int, queue *[]int, to_return *[]int) bool {
	if len((*queue)) == 0 {
		return true
	}
	node := (*queue)[0]

	(*queue) = (*queue)[1:]

	if visited[node] {

		return false
	}
	if len(adj[node]) == 0 {
		if !added[node] {
			(*to_return) = append((*to_return), node)
		}
		added[node] = true

		return true
	}
	visited[node] = true
	(*queue) = append((*queue), adj[node]...)
	if !findorder_bfs(visited, added, adj, queue, to_return) {
		return false
	}
	delete(visited, node)

	// adj[node] = []int{}
	if !added[node] {
		(*to_return) = append((*to_return), node)
	}
	added[node] = true

	return true
}
