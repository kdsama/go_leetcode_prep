package grind75

func CourseSchedule(numCourses int, prerequisites [][]int) bool {

	mapping := map[int][]int{}
	for i := range prerequisites {
		mapping[prerequisites[i][0]] = append(mapping[i], prerequisites[i][1])

	}
	// var dfs func(mapping map[int]int, el int)
	set := map[int]int{}
	var dfs func(el int) bool
	dfs = func(el int) bool {

		if _, ok := set[el]; !ok {
			return false
		}
		if len(mapping[el]) == 0 {
			return true
		}
		set[el] = 1
		for pre := range mapping[el] {
			if ok := dfs(mapping[el][pre]); !ok {
				return false
			}

		}
		delete(set, el)
		mapping[el] = []int{}
		return true

	}
	for el := range prerequisites {
		if ok := dfs(prerequisites[el][0]); !ok {
			return false
		}
	}
	return true
}
