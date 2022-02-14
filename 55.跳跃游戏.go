/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	maxIndex := 0 // 能到达的最远位置
	for idx, val := range nums {
		// 当前 idx 如果能够到达，同时从这个位置跳能超过 maxIndex 重置最远索引
		if maxIndex >= idx && idx+val > maxIndex {
			maxIndex = idx + val
		}
	}
	// 是否能够到达最后
	return maxIndex >= len(nums)-1
}

// @lc code=end

