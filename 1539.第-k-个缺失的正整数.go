/*
 * @lc app=leetcode.cn id=1539 lang=golang
 *
 * [1539] 第 k 个缺失的正整数
 */

// @lc code=start
func arrangeCoins(n int) int {
	return sort.Search(n, func(k int) bool { k++; return k*(k+1) > 2*n })
}

// @lc code=end

