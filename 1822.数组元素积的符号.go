/*
 * @lc app=leetcode.cn id=1822 lang=golang
 *
 * [1822] 数组元素积的符号
 */

// @lc code=start
func arraySign(nums []int) int {
	positive := true
	zeroStatus := false
	for _, num := range nums {
		if num == 0 {
			zeroStatus = true
			break
		} else if num < 0 {
			if positive {
				positive = false
			} else {
				positive = true
			}
		}
	}
	if zeroStatus {
		return 0
	} else if positive {
		return 1
	} else {
		return -1
	}
}

// @lc code=end

