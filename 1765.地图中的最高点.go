/*
 * @lc app=leetcode.cn id=1765 lang=golang
 *
 * [1765] 地图中的最高点
 */

// @lc code=start
func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	res := make([][]int, m)
	
	d := [][]int{}		// 水域队列

	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				d = append(d, []int{i, j})	// 水域格子入队
				res[i][j] = 0
			} else {
				res[i][j] = -1
			}
		}
	}

	// 东西南北
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1},{0, -1}}

	for h := 1; len(d) > 0; h++ {
		for size := len(d); size > 0; size-- {
			info := d[0]	// 水域格子出队
			d = d[1:]		// 原队列修改
			x, y := info[0], info[1]
			for _, di := range dirs {	// 东西南北
				nx, ny := x + di[0], y + di[1]
				// 未越界同时是陆地
				if nx >= 0 && nx < m && ny >= 0 && ny < n && res[nx][ny] == -1 {
					res[nx][ny] = h
					// 当前陆地格子入队，下一次循环以陆地格子为起点东南西北增 1
					d = append(d, []int{nx, ny})
				}
			}
		}
	}
	return res
}
// @lc code=end

