/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	length, length0 := len(matrix), len(matrix[0])
	// 超出边界
	if target < matrix[0][0] || target > matrix[length-1][length0-1] {
		return false
	}
	// 确定行号
	row := -1
	if length == 1 { // 下面前移一位就导致长度为 1 的矩阵无法正常获取行号
		row = 0
	} else {
		for i := 0; i < length-1; i++ {
			if matrix[i][0] <= target && matrix[i+1][0] > target {
				row = i
				break // 及时退出
			}
		}
	}
	// 上方如果结果row==-1，那么数字只能在最后一行
	if row == -1 { // 或者matrix[length-1][0] >= target，这样有查询的时间消耗
		row = length - 1
	}
	// 排序情况下====二分查找
	for left, right := 0, length0-1; left <= right; {
		mid := (left + right) / 2
		one := matrix[row][mid]
		if one == target {
			return true
		} else if one < target {
			// 左边界左右移
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 都没找到====无
	return false
}

// @lc code=end

