/*
 * @lc app=leetcode.cn id=934 lang=golang
 *
 * [934] 最短的桥
 */

// @lc code=start
func shortestBridge(grid [][]int) (res int) {
	q := [][2]int{}
	m, n := len(grid), len(grid[0])
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if !(0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1) {
			return
		}
		grid[x][y] = 0
		q = append(q, [2]int{x, y})
		for _, direction := range directions {
			dfs(x+direction[0], y+direction[1])
		}
	}

	find := false
	for i := 0; i < m; i++ {
		if find {
			break
		}
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				find = true
				dfs(i, j)
				break
			}
		}
	}
	visited := map[[2]int]bool{}
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			point := q[0]
			q = q[1:]
			for _, direction := range directions {
				x, y := point[0]+direction[0], point[1]+direction[1]
				if !(0 <= x && x < m && 0 <= y && y < n) || visited[[2]int{x, y}] {
					continue
				}
				if grid[x][y] == 1 {
					return
				} else {
					q = append(q, [2]int{x, y})
					visited[[2]int{x, y}] = true
				}
			}
		}
		res++
	}
	return
}

// @lc code=end

