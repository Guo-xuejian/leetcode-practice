/*
 * @lc app=leetcode.cn id=1144 lang=golang
 *
 * [1144] 递减元素使数组呈锯齿状
 */

// @lc code=start
func movesToMakeZigzag(nums []int) int {
	resOdd, resEven, pre, after := 0, 0, 0, 0
	length := len(nums)
	for i := 0; i < length; i++ {
		// 需要注意，数组索引奇偶和常规数字不一样，从零开始，索引 1 为实际偶数为
		if i%2 == 0 {
			// 奇数需要大于相邻数字
			// 记住，一定是大于等于，等于也不会符合要求
			if i > 0 && nums[i] >= nums[i-1] {
				pre = nums[i] - nums[i-1] + 1
			} else {
				pre = 0
			}
			if i < length-1 && nums[i] >= nums[i+1] {
				after = nums[i] - nums[i+1] + 1
			} else {
				after = 0
			}
			resOdd += max(pre, after)
		} else {
			// 或者偶数需要大于相邻数字
			if i > 0 && nums[i] >= nums[i-1] {
				pre = nums[i] - nums[i-1] + 1
			} else {
				pre = 0
			}
			if i < length-1 && nums[i] >= nums[i+1] {
				after = nums[i] - nums[i+1] + 1
			} else {
				after = 0
			}
			resEven += max(pre, after)
		}
	}
	return min(resOdd, resEven)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

