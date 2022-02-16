/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) (res int) {
	position := len(nums) - 1
	for position > 0 { // 从后往前找，找到索引 0 处返回步骤数 res 即可
		// 从前到后，找到离 position 最远的索引（满足能够到达）
		for i := 0; i < position; i++ {
			if i+nums[i] >= position { // 能够到达
				// 重置跳点，继续找前一个跳点，此时需要退出，不需要前面的挑点（虽然也能跳但是距离变远）
				position = i
				res++
				break
			}
		}
	}
	return res
}

// @lc code=end

