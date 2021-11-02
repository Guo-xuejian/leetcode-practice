/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	length := len(digits)
	// var temp int
	for i := length - 1; i >= 0; i-- {
		// temp = digits[i] + 1
		if digits[i] == 9 { // 只有 该位为9时加1进位
			digits[i] = 0
			if i == 0 { // 类似 99，999，这样的数字需要多出一位
				digits = append([]int{1}, digits...) // 三个点表示切片
			}
		} else { // 不大于10，则直接返回即可
			digits[i] += 1
			return digits
		}
	}
	return digits
}

// @lc code=end

