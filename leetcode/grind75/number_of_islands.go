package grind75

func NumIslands(grid [][]byte) int {

	count := 0
	// we have to see till where the 1s have gone
	// we also have to map what all elements we have already covered so we dont need to cover it again
	arr := make([][]int, len(grid))
	for i := range arr {
		arr[i] = make([]int, len(grid[0]))
	}
	// if we go through an element , we will mark it 1 here .

	for i := range grid {
		for j := range grid[0] {
			if arr[i][j] != 1 {
				count += bfsNumIslands(&arr, grid, i, j)
			}
		}
	}
	return count
}

func bfsNumIslands(arr *[][]int, grid [][]byte, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	if ((*arr)[i][j]) == 1 {
		return 0
	}
	(*arr)[i][j] = 1
	if grid[i][j] == '0' {
		return 0
	}

	bfsNumIslands(arr, grid, i+1, j)
	bfsNumIslands(arr, grid, i, j+1)
	return 1

}
