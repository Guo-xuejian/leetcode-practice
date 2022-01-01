/*
 * @lc app=leetcode.cn id=2022 lang=golang
 *
 * [2022] 将一维数组转变成二维数组
 */

// @lc code=start
func construct2DArray(original []int, m int, n int) (res [][]int) {
	if len(original) != m*n {
		return
	}
	index := 0
	for i := 0; i < m; i++ {
		temp := make([]int, n)
		for j := 0; j < n; j++ {
			temp[j] = original[index]
			index++
		}
		res = append(res, temp)
	}
	return
}

// @lc code=end

