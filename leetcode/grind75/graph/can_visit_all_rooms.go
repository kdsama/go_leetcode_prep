package graph

func CanVisitAllRooms(rooms [][]int) bool {

	mapping := map[int]bool{}
	dfs(rooms, 0, mapping)

	return len(mapping) == len(rooms)
}

func dfs(rooms [][]int, val int, mapping map[int]bool) {

	if _, ok := mapping[val]; ok {
		return
	}
	mapping[val] = true

	for _, v := range rooms[val] {
		dfs(rooms, v, mapping)
	}

}
