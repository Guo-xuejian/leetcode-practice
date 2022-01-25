/*
 * @lc app=leetcode.cn id=1688 lang=golang
 *
 * [1688] 比赛中的配对次数
 */

// @lc code=start
func numberOfMatches(n int) (res int) {
	for n > 1 {
		res += n / 2
		if n % 2 != 0 {
			n = n / 2 + 1
		} else {
			n = n / 2
		}
	}
	return
}
// @lc code=end

