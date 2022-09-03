/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */

// @lc code=start
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })
	n := len(pairs)
	dp := make([]int, n)
	for i, p := range pairs {
		dp[i] = 1
		for j, q := range pairs[:i] {
			if p[0] > q[1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

