/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */

// @lc code=start
func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, 0, m*n)
	for i := 0; i < m+n-1; i++ {
		if i%2 == 1 {
			x := max(i-n+1, 0)
			y := min(i, n-1)
			for x < m && y >= 0 {
				ans = append(ans, mat[x][y])
				x++
				y--
			}
		} else {
			x := min(i, m-1)
			y := max(i-m+1, 0)
			for x >= 0 && y < n {
				ans = append(ans, mat[x][y])
				x--
				y++
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

