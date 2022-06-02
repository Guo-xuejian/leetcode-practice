/*
 * @lc app=leetcode.cn id=829 lang=golang
 *
 * [829] 连续整数求和
 */

// @lc code=start
func isKConsecutive(n, k int) bool {
	if k%2 == 1 {
		return n%k == 0
	}
	return n%k != 0 && 2*n%k == 0
}

func consecutiveNumbersSum(n int) (ans int) {
	for k := 1; k*(k+1) <= n*2; k++ {
		if isKConsecutive(n, k) {
			ans++
		}
	}
	return
}

// @lc code=end

