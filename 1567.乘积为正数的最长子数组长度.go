/*
 * @lc app=leetcode.cn id=1567 lang=golang
 *
 * [1567] 乘积为正数的最长子数组长度
 */

// @lc code=start
func getMaxLen(nums []int) (res int) {
	m := len(nums)
	positive, negative := make([]int, m), make([]int, m)
	if nums[0] > 0 {
		positive[0] = 1
	} else if nums[0] < 0 {
		negative[0] = 1
	}
	res = positive[0]
	for i := 1; i < m; i++ {
		if nums[i] > 0 {
			positive[i] = positive[i-1] + 1
			if negative[i-1] > 0 {
				negative[i] = negative[i-1] + 1
			} else {
				negative[i] = 0
			}
		} else if nums[i] < 0 {
			if negative[i-1] > 0 {
				positive[i] = negative[i-1] + 1
			} else {
				positive[i] = 0
			}
			negative[i] = positive[i-1] + 1
		} else {
			positive[i] = 0
			negative[i] = 0
		}
		res = max(res, positive[i])
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

