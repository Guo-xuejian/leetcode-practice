/*
 * @lc app=leetcode.cn id=1572 lang=golang
 *
 * [1572] 矩阵对角线元素的和
 */

// @lc code=start
func diagonalSum(mat [][]int) (res int) {
	m := len(mat)
	index := 0
	for index < m {
		res += mat[index][index] + mat[index][m-1-index] + mat[m-1-index][index] + mat[m-1-index][m-1-index]
		index++
	}
	// 存在重复计数，对角线元素 2 次，中央元素（如果存在）总计 4 次
	res /= 2      // 去重，中央仍多一次
	if m%2 != 0 { // 减去存在的重要元素
		res -= mat[m/2][m/2]
	}
	return
}

// @lc code=end

