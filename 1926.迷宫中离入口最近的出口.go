/*
 * @lc app=leetcode.cn id=1926 lang=golang
 *
 * [1926] 迷宫中离入口最近的出口
 */

// @lc code=start
func nearestExit(maze [][]byte, entrance []int) int {
	// 直接 bfs
	m, n := len(maze), len(maze[0])
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	q := [][2]int{{entrance[0], entrance[1]}}

	// 判定是否越界
	inArea := func(x, y int) bool {
		if 0 <= x && x < m && 0 <= y && y < n {
			return true
		}
		return false
	}
	maze[entrance[0]][entrance[1]] = '+'
	for res := 1; len(q) > 0; res++ {
		length := len(q)
		for i := 0; i < length; i++ {
			point := q[0]
			q = q[1:]
			for _, direction := range directions {
				x, y := point[0]+direction[0], point[1]+direction[1]
				if inArea(x, y) && maze[x][y] == '.' {
					if x == 0 || y == 0 || x == m-1 || y == n-1 {
						return res
					}
					maze[x][y] = '+'
					q = append(q, [2]int{x, y})
				}
			}
		}
	}
	// 还没找到，置为 -1 返回
	return -1
}
// @lc code=end

