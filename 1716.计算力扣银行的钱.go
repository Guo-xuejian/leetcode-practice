/*
 * @lc app=leetcode.cn id=1716 lang=golang
 *
 * [1716] 计算力扣银行的钱
 */

// @lc code=start
func totalMoney(n int) (res int) {
	currDay := 0 // 模拟一个第零天，后续相加即可
	for i := 0; i < n; i++ {
		currDay++
		if i%7 == 0 && i != 0 {
			currDay -= 6 // currDay = currDay - 7 + 1
		}
		res += currDay
	}
	return res
}

// @lc code=end

