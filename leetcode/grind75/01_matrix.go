package grind75

func UpdateMatrix(mat [][]int) [][]int {
	// if we add them all we get prefix sums
	//  0 0 0 0 1 0 1 1 1
	// I think some kind of recursion would be here //
	// l := len(mat)
	// m := len(mat[0])
	new_matrix := [][]int{}
	for i := range mat {
		new_matrix = append(new_matrix, []int{})
		for j := range mat[i] {
			if j == 1 {
				// println(mat[i][j])
			}
			if i == 2 {
				println("HI ", i, j, mat[i][j])
			}
			new_matrix[i] = append(new_matrix[i], getDistance(mat, i, j))

		}

	}
	return new_matrix
}

func min(a int, b int, c int, d int) int {
	min := 2
	j := []int{a, b, c, d}
	for i := range j {
		if j[i] < min {
			min = j[i]
		}
	}
	return min
}

func getDistance(mat [][]int, i int, j int) int {
	m := len(mat)
	n := len(mat[0])

	a, b, c, d := 0, 0, 0, 0
	if i+1 < m && j+1 < n && i > 0 && j > 0 {

		if mat[i][j] == 0 {
			return 0
		}

		b = 1 + getDistance(mat, i+1, j)

		a = 1 + getDistance(mat, i-1, j)

		c = 1 + getDistance(mat, i, j+1)

		d = 1 + getDistance(mat, i, j-1)
	}

	return min(a, b, c, d)

}
