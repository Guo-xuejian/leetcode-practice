/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) (res int) {
	// 将所有属于同一个岛屿的陆地格子全部更改为 2，后续遍历不再查看该格子
	m, n := len(grid), len(grid[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		// 先判定索引越界，后续遍历陆地格子
		if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] != '1' {
			return
		}
		grid[x][y] = '2'
		for _, direction := range directions {
			dfs(x+direction[0], y+direction[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' { // 遍历每一个还没有被访问的陆地格子
				dfs(i, j)
				res++ // 只有 dfs() 外部可以增加结果，dfs()只做各个相邻陆地格子的访问和修改
			}
		}
	}
	return
}

// @lc code=end

