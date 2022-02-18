/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */

// @lc code=start
func pacificAtlantic(heights [][]int) (res [][]int) {
	m, n := len(heights), len(heights[0])
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	altanicOcean := make([][]int, m)
	pacificOcean := make([][]int, m)
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		altanicOcean[i] = make([]int, n)
		pacificOcean[i] = make([]int, n)
		visited[i] = make([]int, n)
	}
	var inArea func(x, y int) bool
	var dfs func(x, y int, flag bool)
	inArea = func(x, y int) bool {
		if 0 <= x && x < m && 0 <= y && y < n {
			return true
		}
		return false
	}
	dfs = func(x, y int, flag bool) {
		if visited[x][y] == 1 {
			return
		}
		visited[x][y] = 1
		if flag {
			pacificOcean[x][y] = 1
		} else {
			altanicOcean[x][y] = 1
		}

		for i := 0; i < 4; i++ {
			currX, currY := x+directions[i][0], y+directions[i][1]
			if !inArea(currX, currY) || heights[x][y] > heights[currX][currY] {
				continue
			}
			dfs(currX, currY, flag)
		}
		return
	}

	// 从各面逐步攀登至顶峰，比较从大西洋和太平洋dfs的山峰是否存在相等坐标
	// 上太平洋
	for i := 0; i < n; i++ {
		dfs(0, i, true)
	}
	visited = make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0, true)
	}
	visited = make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dfs(m-1, j, false)
	}
	visited = make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	for j := 0; j < m; j++ {
		dfs(j, n-1, false)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacificOcean[i][j] == 1 && altanicOcean[i][j] == 1 {
				res = append(res, []int{i, j})
			}
		}
	}
	return
}

// @lc code=end

