/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 */

// @lc code=start
func dominantIndex(nums []int) int {
	// 通过一轮便利维护一个最大值和一个次大值，如果最大值至少是次大值的两倍，那么自然满足条件
	maxNum, preMax, index := nums[0], -1, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum, preMax = nums[i], maxNum
			index = i
		} else if nums[i] < maxNum && nums[i] > preMax {
			preMax = nums[i]
		}
	}
	// 判断是否至少是两倍
	if maxNum >= preMax*2 {
		return index
	} else {
		return -1
	}
}

// @lc code=end

