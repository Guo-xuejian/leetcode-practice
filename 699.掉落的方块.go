/*
 * @lc app=leetcode.cn id=699 lang=golang
 *
 * [699] 掉落的方块
 */

// @lc code=start
func fallingSquares(positions [][]int) []int {
	n := len(positions)
	heights := make([]int, n)
	for i, p := range positions {
		left1, right1 := p[0], p[0]+p[1]-1
		heights[i] = p[1]
		for j, q := range positions[:i] {
			left2, right2 := q[0], q[0]+q[1]-1
			if right1 >= left2 && right2 >= left1 {
				heights[i] = max(heights[i], heights[j]+p[1])
			}
		}
	}
	for i := 1; i < n; i++ {
		heights[i] = max(heights[i], heights[i-1])
	}
	return heights
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

