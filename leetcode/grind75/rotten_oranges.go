package grind75

func orangesRotting(grid [][]int) int {
	// use a queue
	// push all the rottentomatoes to the queue
	// pop
	// check the surroundings
	// if a  normal tomato , normal tomato--
	// time++ after checking all four directions

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	r, c := len(grid), len(grid[0])
	queue := [][]int{}
	rt := 0
	time := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				rt++
			}
		}
	}

	for len(queue) > 0 && rt > 0 {

		ll := len(queue)

		for i := 0; i < ll; i++ {
			a, b := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, dir := range directions {
				x := a + dir[0]
				y := b + dir[1]
				if (x < 0 || y < 0 || x >= r || y >= c) || grid[x][y] != 1 {
					continue
				}
				grid[x][y] = 2
				queue = append(queue, []int{x, y})
				rt--
			}
		}
		time++
	}
	if rt > 0 {
		return -1
	}
	return time
}
