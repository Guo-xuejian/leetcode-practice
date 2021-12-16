/*
 * @lc app=leetcode.cn id=1014 lang=golang
 *
 * [1014] 最佳观光组合
 */

// @lc code=start
func maxScoreSightseeingPair(values []int) int {
	maxScore, mx := 0, values[0]+0
	for j := 1; j < len(values); j++ {
		maxScore = max(maxScore, mx+values[j]-j)
		// 一边遍历一边维护前方的景点，取前方景点中值与索引之和最大的
		mx = max(mx, values[j]+j)
	}
	return maxScore
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

