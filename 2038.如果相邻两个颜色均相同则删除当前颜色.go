/*
 * @lc app=leetcode.cn id=2038 lang=golang
 *
 * [2038] 如果相邻两个颜色均相同则删除当前颜色
 */

// @lc code=start
func winnerOfGame(colors string) bool {
	var res int
	for i := 1; i < len(colors)-1; i++ {
		if colors[i] == colors[i-1] && colors[i] == colors[i+1] {
			if colors[i] == 'A' {
				res += 1
			} else {
				res -= 1
			}
		}
	}
	return res > 0
}

// @lc code=end

