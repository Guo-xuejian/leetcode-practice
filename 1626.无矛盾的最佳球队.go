/*
 * @lc app=leetcode.cn id=1626 lang=golang
 *
 * [1626] 无矛盾的最佳球队
 */

// @lc code=start
func bestTeamScore(scores []int, ages []int) (res int) {
	length := len(ages)
	ageSortedScores := make([][2]int, length)
	for idx, val := range ages {
		ageSortedScores[idx] = [2]int{val, scores[idx]}
	}
	ageSortedScores = append([][2]int{{0, 0}}, ageSortedScores...) // 添加一个哨兵
	sort.Slice(ageSortedScores, func(i, j int) bool {
		if ageSortedScores[i][0] == ageSortedScores[j][0] {
			return ageSortedScores[i][1] < ageSortedScores[j][1]
		}
		return ageSortedScores[i][0] < ageSortedScores[j][0]
	})
	dp := make([]int, length+1)
	for i := 1; i < length+1; i++ {
		for j := 0; j < i; j++ {
			// 遍历所有同龄或者低龄者，需满足分数小于等于当前球员
			if ageSortedScores[i][1] >= ageSortedScores[j][1] {
				// 获取最长增长子序列
				dp[i] = max(dp[i], ageSortedScores[i][1]+dp[j])
				res = max(res, dp[i])
			}
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

