/*
 * @lc app=leetcode.cn id=566 lang=golang
 *
 * [566] 重塑矩阵
 */

// @lc code=start
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if r*c != m*n {
		return mat
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	for x := 0; x < m*n; x++ {
		res[x/c][x%c] = mat[x/n][x%n]
	}
	return res
}

// @lc code=end

