/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(mat [][]int) [][]int {
	n, m := len(mat), len(mat[0])
	q := make([][]int, 0)
	// 把0全部存进队列，后面从队列中取出来，判断每个访问过的节点的上下左右，直到所有的节点都被访问过为止。
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(q) > 0 {
		point := q[0]
		q = q[1:]
		for _, v := range direction {
			x := point[0] + v[0]
			y := point[1] + v[1]
			// 越界判定
			if x >= 0 && x < n && y >= 0 && y < m && mat[x][y] == -1 {
				// 为什么不直接 0 + 1，因为 1 可能被 1 包围，所以外层 --> 内层，外层数为内层数的基础
				// 若仍然存在内部 0 ，则外层的修改会被覆盖为最小
				mat[x][y] = mat[point[0]][point[1]] + 1
				q = append(q, []int{x, y}) // 写入当前，作为与其相邻的 1 的基础
			}
		}
	}
	return mat
}

// @lc code=end

