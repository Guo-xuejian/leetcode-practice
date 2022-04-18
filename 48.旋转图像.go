/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	temp := make([][]int, n)
	for i := 0; i < n; i++ {
		temp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			temp[j][n-i-1] = matrix[i][j]
		}
	}
	copy(matrix, temp)
}

// @lc code=end

