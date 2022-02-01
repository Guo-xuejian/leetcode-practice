/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */

// @lc code=start
func superEggDrop(k int, n int) int {
	if k < 1 || n < 1 {
		return 0
	}
	dp := make([]int, k)
	res := 0
	for {
		res++
		pre := 0
		for i := 0; i < len(dp); i++ {
			temp := dp[i]
			dp[i] = dp[i] + pre + 1
			pre = temp
			if dp[i] >= n {
				return res
			}
		}
	}
}
// @lc code=end

