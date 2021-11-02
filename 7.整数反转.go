/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	res := 0
	for x != 0 {
		// 不在 [−2**31,  2**31 − 1] 范围之内，则返回0，这里的**只是代表次方，不能代码实现
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		// 处在合理范围，% 获取此次低进位数字
		digit := x % 10
		// 其余高进位保留，并 / 10 去除这一次的低进位
		x /= 10
		// 第一次 res == 0，所以第一次获得代表最终数字最高位的数字在这次计算中没有 *10 ，后续就会正常，啊哈有意思
		res = res*10 + digit
	}
	return res
}

// @lc code=end

