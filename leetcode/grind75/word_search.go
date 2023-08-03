package grind75

func exist(board [][]byte, word string) bool {
	// we need to search for a word on the board
	// once a letter is marked we should be going back to it unless the word fails

	boundary := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(i, j int, str string) bool

	rows := len(board)
	col := len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, col)
	}

	dfs = func(a, b int, str string) bool {

		if len(str) == 0 {
			return true
		}
		// fmt.Println("looking for ",string(str[0]),"and we are at previous character found at ",a,b)
		// we have an array ab <-- so it must be sent as well
		// find the first character of the word
		// we need to pass i,j
		// make sure that for scenarios like ABA we dont go back to previous position

		// issue for my case is
		for _, val := range boundary {
			i, j := a+val[0], b+val[1]

			if i < 0 || i >= rows || j < 0 || j >= col || visited[i][j] || board[i][j] != str[0] {
				continue
			}
			visited[i][j] = true
			ans := dfs(i, j, str[1:])

			if ans {

				return ans
			}
			visited[i][j] = false
		}

		return false

	}
	for i := 0; i < rows; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] {
				visited[i][j] = true
				res := dfs(i, j, word[1:])
				if res {
					return res
				}
				visited[i][j] = false
			}
		}
	}
	return false

}
