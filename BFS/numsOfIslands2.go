package bfs

type unionFind struct {
	parent []int
	rank   []int
}

func (u *unionFind) init(n int) {
	u.parent = make([]int, n)
	u.rank = make([]int, n)
	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.rank[i] = 1
	}
}
func (u *unionFind) find(x int) int {
	if u.parent[x] == x {
		return x
	}
	u.parent[x] = u.find(u.parent[x])
	return u.parent[x]
}

func (u *unionFind) union(x, y int) {
	rx, ry := u.find(x), u.find(y)
	if rx == ry {
		return
	}
	if u.rank[rx] < u.rank[ry] {
		u.parent[rx] = ry
	} else if u.rank[rx] > u.rank[ry] {
		u.parent[ry] = rx
	} else {
		u.parent[ry] = rx
		u.rank[rx]++
	}
}

func numIslands2(grid [][]byte) int {

	res := 0
	m := len(grid)
	if m == 0 {
		return res
	}

	n := len(grid[0])
	dsu := unionFind{}
	dsu.init(m * n)
	move := []direction{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				for k := 0; k < len(move); k++ {
					x, y := i+move[k].leftRight, j+move[k].upDown
					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1' {
						if dsu.find(i*n+j) != dsu.find(x*n+y) {
							res--
						}
						dsu.union(i*n+j, x*n+y)
					}
				}
			}
		}
	}
	return res
}
