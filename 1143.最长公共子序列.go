/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) (maxLength int) {
	text1Length := len(text1)
	text2Length := len(text2)
	dp := make([][]int, text1Length+1)
	for i := range dp {
		dp[i] = make([]int, text2Length+1)
	}
	for i, v1 := range text1 {
		for j, v2 := range text2 {
			if v1 == v2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[text1Length][text2Length]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

