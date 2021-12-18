/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) (currentRow []int) {
	preRow := []int{}
	// 由于只和上一层有关，所以实际上一需要维护当前行和上一行
	for i := 0; i < rowIndex+1; i++ {
		currentRow = []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				currentRow = append(currentRow, 1)
			} else {
				currentRow = append(currentRow, preRow[j]+preRow[j-1])
			}
		}
		// if i == rowIndex { // 索引符合条件即返回当前行
		// 	return currentRow
		// } else { // 不符合条件则重置前一行
		// 	preRow = currentRow
		// }
		preRow = currentRow
	}
	return preRow
}

// @lc code=end

