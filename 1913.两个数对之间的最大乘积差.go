/*
 * @lc app=leetcode.cn id=1913 lang=golang
 *
 * [1913] 两个数对之间的最大乘积差
 */

// @lc code=start
func maxProductDifference(nums []int) int {
	// 找出最大值和次大值，最小值和次小值
	maxIndex, preMaxIndex, minIndex, preMinIndex := 0, 0, 0, 0
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		// 最小值和次小值
		if curr <= nums[minIndex] {
			minIndex, preMinIndex = i, minIndex
		} else if curr > nums[minIndex] && (curr < nums[preMinIndex] || preMinIndex == 0 && minIndex == 0) {
			preMinIndex = i
		}

		// 最大值和次大值
		if curr >= nums[maxIndex] {
			maxIndex, preMaxIndex = i, maxIndex
		} else if curr < nums[maxIndex] && (curr > nums[preMaxIndex] || preMaxIndex == 0 && maxIndex == 0) {
			preMaxIndex = i
		}
	}
	return nums[maxIndex]*nums[preMaxIndex] - nums[minIndex]*nums[preMinIndex]
}

// @lc code=end

