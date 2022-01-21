/*
 * @lc app=leetcode.cn id=848 lang=golang
 *
 * [848] 字母移位
 */

// @lc code=start
func shiftingLetters(s string, shifts []int) string {
	res := make([]byte, len(s))
	shifts = append(shifts, 0) // 方便后续处理，添加一个
	for i := len(s) - 1; i >= 0; i-- {
		// 循环内修改数字,单开循环会超时
		shifts[i] = (shifts[i] + shifts[i+1]) % 26
		// 对应字符移位写入 res, 需要注意是个循环，所以需要取余
		res[i] = 'a' + (s[i]+byte(shifts[i])-'a')%26
	}
	// 转换为字符串返回
	return string(res)
}

// @lc code=end

