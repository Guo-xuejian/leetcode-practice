/*
 * @lc app=leetcode.cn id=479 lang=golang
 *
 * [479] 最大回文数乘积
 */

// @lc code=start
func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	upper := int(math.Pow10(n)) - 1
	for left := upper; ; left-- { // 枚举回文数的左半部分
		p := left
		for x := left; x > 0; x /= 10 {
			p = p*10 + x%10 // 翻转左半部分到其自身末尾，构造回文数 p
		}
		for x := upper; x*x >= p; x-- {
			if p%x == 0 { // x 是 p 的因子
				return p % 1337
			}
		}
	}
}

// @lc code=end

