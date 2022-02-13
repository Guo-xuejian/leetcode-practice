/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	m, n := len(image), len(image[0])
	preColor := image[sr][sc]
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x >= m || y >= n {
			return
		}
		if image[x][y] == preColor {
			image[x][y] = newColor
			for _, direction := range directions {
				dfs(x+direction[0], y+direction[1])
			}
		}
	}

	if preColor != newColor {
		dfs(sr, sc)
	}
	return image
}

// @lc code=end

