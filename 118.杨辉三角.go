/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) (res [][]int) {
	for i := 0; i < numRows; i++ {
		row := []int{} // 初始化本行
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i { // 边界则加入 1
				row = append(row, 1)
			} else { // 其余情况按照要求加入左上方与右上方之和
				row = append(row, res[i-1][j]+res[i-1][j-1])
			}
		}
		res = append(res, row)
	}
	return
}

// @lc code=end

