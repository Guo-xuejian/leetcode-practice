/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除并获得点数
 */

// @lc code=start
func deleteAndEarn(nums []int) int {
	// 和打家劫舍一样，不能偷取相邻的,但是这里有一些区别，nums[i] - 1 和 nums[i] + 1 可能不存在
	// 那就找到最大值 maxNum 构建一个 [1,2,3,4,.....maxNum]
	if len(nums) == 1 {
		return nums[0]
	}
	maxNum := 0
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}
	dp := make([]int, maxNum+1)
	for _, num := range nums {
		dp[num] += num // 相同数字不要紧，一起偷了就行
	}
	return rob(dp)
}

func rob(nums []int) int {
	// 打家劫舍
	preNum, currNum := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		preNum, currNum = currNum, max(preNum+nums[i], currNum)
	}
	return currNum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

