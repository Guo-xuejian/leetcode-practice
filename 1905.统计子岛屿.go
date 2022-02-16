/*
 * @lc app=leetcode.cn id=1905 lang=golang
 *
 * [1905] 统计子岛屿
 */

// @lc code=start
func countSubIslands(grid1 [][]int, grid2 [][]int) (res int) {
	m, n := len(grid1), len(grid1[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		// 内部递归时只处理陆地格子，不能越界
		if !(x < m && x >= 0 && y < n && y >= 0 && grid2[x][y] == 1) {
			return
		}
		grid2[x][y] = 0 // 剪枝

		for _, direction := range directions {
			dfs(x+direction[0], y+direction[1])
		}
	}
	// 清除不符合条件的岛屿
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				dfs(i, j)
			}
		}
	}
	// 计算子岛屿
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				dfs(i, j)
			}
		}
	}
	return
}

// @lc code=end

