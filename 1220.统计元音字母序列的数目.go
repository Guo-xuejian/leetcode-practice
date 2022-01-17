/*
 * @lc app=leetcode.cn id=1220 lang=golang
 *
 * [1220] 统计元音字母序列的数目
 */

// @lc code=start
func countVowelPermutation(n int) (res int) {
	const mod int = 1e9 + 7
	dp := []int{1, 1, 1, 1, 1}
	for i := 0; i < n-1; i++ {
		aEnd, eEnd, iEnd, oEnd, uEnd := dp[0], dp[1], dp[2], dp[3], dp[4]
		dp[0] = (eEnd + iEnd + uEnd) % mod // a 结尾的就是 e(1) i(2) u(4)
		dp[1] = (aEnd + iEnd) % mod        // e 结尾的就是 a(0) i(2)
		dp[2] = (eEnd + oEnd) % mod        // i 结尾的就是 e(1) o(3)
		dp[3] = iEnd % mod                 // o 结尾的就是 i(2)
		dp[4] = (iEnd + oEnd) % mod        // u 结尾的就是 i(2) o(3)
	}
	return sum(dp...) % mod
}

func sum(a ...int) (res int) {
	for _, one := range a {
		res += one
	}
	return
}

// @lc code=end

