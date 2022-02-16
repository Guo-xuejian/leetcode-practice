/*
 * @lc app=leetcode.cn id=918 lang=golang
 *
 * [918] 环形子数组的最大和
 */

// @lc code=start
func maxSubarraySumCircular(nums []int) (res int) {
	total, maxSum, currMax, minSum, currMin := 0, nums[0], 0, nums[0], 0

	for _, num := range nums {
		total += num
		// 维护递增
		currMax = max(currMax+num, num)
		maxSum = max(currMax, maxSum)
		// 维护递减
		currMin = min(currMin+num, num)
		minSum = min(currMin, minSum)
	}
	if maxSum > 0 {
		res = max(maxSum, total-minSum)
	} else {
		res = maxSum
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

