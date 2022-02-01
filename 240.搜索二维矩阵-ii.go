/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
    length1, length2 := len(matrix), len(matrix[0])

	for i := 0; i < length1; i++ {
		// 当前行最小值仍大于，同时大于上一行最大值，无法满足
		if matrix[i][0] > target {
			return false
		}
		// 在当前行，二分查找
		if matrix[i][length2-1] >= target {
			left, right := 0, length2 - 1
			for left <= right {
				mid := (left + right) / 2
				curr := matrix[i][mid]
				if curr == target {
					return true
				} else if curr > target {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
	}
	// 遍历完毕未找到
	return false
}
// @lc code=end

