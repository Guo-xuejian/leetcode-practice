/*
 * @lc app=leetcode.cn id=1252 lang=golang
 *
 * [1252] 奇数值单元格的数目
 */

// @lc code=start
func oddCells(m int, n int, indices [][]int) (res int) {
	// 模拟过程最后计算奇数数字数量
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for _, val := range indices {
		for i := 0; i < m; i++ {
			matrix[i][val[1]]++
		}
		for j := 0; j < n; j++ {
			matrix[val[0]][j]++
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j]%2 == 1 {
				res++
			}
		}
	}
	return
}

// 20223-07-12
// func oddCells(m, n int, indices [][]int) (ans int) {
//     matrix := make([][]int, m)
//     for i := range matrix {
//         matrix[i] = make([]int, n)
//     }
//     for _, p := range indices {
//         for j := range matrix[p[0]] {
//             matrix[p[0]][j]++
//         }
//         for _, row := range matrix {
//             row[p[1]]++
//         }
//     }
//     for _, row := range matrix {
//         for _, v := range row {
//             ans += v % 2
//         }
//     }
//     return
// }

// @lc code=end

