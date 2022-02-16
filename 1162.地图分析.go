/*
 * @lc app=leetcode.cn id=1162 lang=golang
 *
 * [1162] 地图分析
 */

// @lc code=start
type Point struct {
	X int
	Y int
}

func maxDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var queue []*Point
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, &Point{i, j})
			}
		}
	}
	if len(queue) == 0 || len(queue) == m*n {
		return -1
	}

	res := 0
	d := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) > 0 {
		res++
		tmpQueu := queue
		queue = nil
		for len(tmpQueu) > 0 {
			p := tmpQueu[0]
			tmpQueu = tmpQueu[1:]
			// 以p为原点，检查4个方向
			for i := 0; i < 4; i++ {
				x := p.X + d[i][0]
				y := p.Y + d[i][1]
				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != 0 {
					continue
				}
				queue = append(queue, &Point{x, y})
				grid[x][y] = 2 // 代表以及被遍历过了
			}
		}
	}

	return res - 1
}

// @lc code=end

