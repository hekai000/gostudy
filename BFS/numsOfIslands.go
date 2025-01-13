package bfs

func numIslands(grid [][]byte) int {
	//dfs
	res := 0
	m := len(grid)
	if m == 0 {
		return res
	}

	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				//dfs(grid, i, j)
				bfs(grid, i, j)
			}
		}
	}
	return res
}

type position struct {
	rowIndex int
	colIndex int
}

type direction struct {
	leftRight int
	upDown    int
}

func bfs(grid [][]byte, row int, col int) {

	q := make([]position, 0)
	move := []direction{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	q = append(q, position{row, col})
	grid[row][col] = '0'
	for len(q) > 0 {

		q = q[1:]

		for i := 0; i < len(move); i++ {
			newX := row + move[i].leftRight
			newY := col + move[i].upDown
			if (newX >= 0) && (newX < len(grid)) && (newY >= 0) && (newY < len(grid[0])) && (grid[newX][newY] == '1') {
				grid[newX][newY] = '0'
				q = append(q, position{newX, newY})
			}

		}
	}

}
