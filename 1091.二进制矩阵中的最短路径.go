/*
 * @lc app=leetcode.cn id=1091 lang=golang
 *
 * [1091] 二进制矩阵中的最短路径
 */

// @lc code=start
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] != 0 {
		return -1
	}
	m := len(grid)
	if m == 1 {
		return 1
	}
	directions := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	q := [][2]int{{0, 0}}
	visited := map[[2]int]bool{[2]int{0, 0}: true}
	start := 1
	inArea := func(x, y int) bool {
        // 不能越界，下一个值为 0，且还没有被访问
		if 0 <= x && 0 <= y && x < m && y < m && grid[x][y] == 0 && !visited[[2]int{x, y}] {
			return true
		}
		return false
	}
	for len(q) > 0 {
		currLength := len(q)
		for i := 0; i < currLength; i++ {
			x , y := q[len(q) - 1][0], q[len(q) - 1][1]
			q = q[:len(q) - 1]   // 覆盖
			for _, direction := range directions {
				currX, currY := x + direction[0], y + direction[1]
				if inArea(currX, currY) {
					if currX == m - 1 && currY == m - 1 {
                        // 已经到终点，加上当前格子退出
						return start + 1
					}
					q = append([][2]int{{currX, currY}}, q...)  // 添加至队首
                    visited[[2]int{currX, currY}] = true    // 已访问
				}
			}
		}
        start++     // 正常 +1
	}
    // 找不到
    return -1
}
// @lc code=end

