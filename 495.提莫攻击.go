/*
 * @lc app=leetcode.cn id=495 lang=golang
 *
 * [495] 提莫攻击
 */

// @lc code=start
func findPoisonedDuration(timeSeries []int, duration int) (res int) {
	poinousTime := 0
	for _, attackTime := range timeSeries {
		if attackTime >= poinousTime {
			res += duration
		} else {
			res += attackTime + duration - poinousTime
		}
		poinousTime = attackTime + duration
	}
	return
}
// @lc code=end

