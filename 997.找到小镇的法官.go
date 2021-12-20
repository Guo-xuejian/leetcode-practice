/*
 * @lc app=leetcode.cn id=997 lang=golang
 *
 * [997] 找到小镇的法官
 */

// @lc code=start
func findJudge(n int, trust [][]int) int {
	res := make([]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = 0
	}
	for _, v := range trust {
		res[v[1]-1]++ // 被相信者
		res[v[0]-1]-- // 相信者
	}
	for i := 0; i < n; i++ {
		if res[i] == n-1 { // 所有人都相信者
			return i + 1
		}
	}
	return -1 // 没有
}

// @lc code=end

