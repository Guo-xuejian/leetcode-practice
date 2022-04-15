// 剑指 Offer 04. 二维数组中的查找
// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

// 示例:

// 现有矩阵 matrix 如下：

// [
//   [1,   4,  7, 11, 15],
//   [2,   5,  8, 12, 19],
//   [3,   6,  9, 16, 22],
//   [10, 13, 14, 17, 24],
//   [18, 21, 23, 26, 30]
// ]
// 给定 target = 5，返回 true。

// 给定 target = 20，返回 false。

// 限制：

// 0 <= n <= 1000

// 0 <= m <= 1000

func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	if matrix[0][0] > target || target > matrix[m-1][n-1] {
		return false
	}
	for i := 0; i < m; i++ {
		// 判断边界是否满足大小关系
		if matrix[i][0] <= target && matrix[i][n-1] >= target {
			// 满足进行二分查找
			left, right := 0, n-1
			for left <= right {
				mid := (left + right) / 2
				if matrix[i][mid] == target {
					return true
				} else if matrix[i][mid] > target {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
	}
	// 没找到
	return false
}

// 2022-04-16
// func findNumberIn2DArray(matrix [][]int, target int) bool {
//     if len(matrix) == 0 || len(matrix[0]) == 0{
//         return false
//     }
//     m, n := len(matrix), len(matrix[0])
//     for i := 0; i < m; i++ {
//         if matrix[i][0] <= target && matrix[i][n - 1] >= target {
//             // 此处开始二分查找
//             left, right := 0, n - 1
//             for left <= right {
//                 mid := left + (right - left) / 2
//                 curr := matrix[i][mid]
//                 if curr == target {
//                     return true
//                 } else if curr > target {
//                     right = mid - 1
//                 } else {
//                     left = mid + 1
//                 }
//             }
//         } else if matrix[i][0] > target {
//             return false
//         } else if matrix[i][n - 1] < target {
//             continue
//         }
//     }
//     return false
// }