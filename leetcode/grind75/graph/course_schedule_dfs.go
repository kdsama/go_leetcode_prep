package graph

func CourseSchedule(numCourses int, prerequisites [][]int) bool {

	visited := map[int]bool{}

	adjList := map[int][]int{}
	for i := range prerequisites {
		adjList[prerequisites[i][0]] = append(adjList[prerequisites[i][0]], prerequisites[i][1])
	}

	for key, _ := range adjList {
		if !visited[key] {
			if !CourseScheduleDFS(&adjList, visited, key) {
				return false
			}
		}
	}
	return true
}

func CourseScheduleDFS(adjList *map[int][]int, visited map[int]bool, key int) bool {
	to_return := true
	if visited[key] {
		return false
	}
	if len((*adjList)[key]) == 0 {
		return true
	}
	visited[key] = true
	for i := range (*adjList)[key] {
		if !CourseScheduleDFS(adjList, visited, (*adjList)[key][i]) {
			return false
		}
	}
	delete(visited, key)
	(*adjList)[key] = []int{}
	return to_return
}
