/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) (ans string) {
	if num1 == "0" || num2 == "0" { // 存在零值直接返回
		return "0"
	}
	m, n := len(num1), len(num2)
	ansArr := make([]int, m+n) // 最终数字长度不会超过两数长度之和
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'       // 码值差异计算数字
		for j := n - 1; j >= 0; j-- { // 一个数的低位与另一个数的较高位逐一乘积
			y := int(num2[j]) - '0'
			ansArr[i+j+1] += x * y
		}
	}
	// 相乘很可能出现两位数
	for i := m + n - 1; i > 0; i-- {
		ansArr[i-1] += ansArr[i] / 10 // 高位进 1，因为不会出现三位数
		ansArr[i] %= 10               // 低位取余
	}
	idx := 0            // 当 1*1 = 1 时 ansArr 会出现首位无值（但是有切片的初始化值 0）
	if ansArr[0] == 0 { // 为 0 时取 1，因为长度至少是 m+n-1
		idx = 1
	}
	for ; idx < m+n; idx++ { // 逐位添加位字符串
		ans += strconv.Itoa(ansArr[idx])
	}
	return ans
}

// @lc code=end

