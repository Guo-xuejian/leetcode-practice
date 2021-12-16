/*
 * @lc app=leetcode.cn id=1936 lang=golang
 *
 * [1936] 新增的最少台阶数
 */

// @lc code=start
func addRungs(rungs []int, dist int) int {
	// 严格递增，计算每次的高度差和 dist 取商就是当前这一步所需要的梯子数目
	stairsNum := 0
	currentHeight := 0
	for _, height := range rungs {
		// 高度要高于 dist 才需要梯子，所以减一，只有 dist+1 才需要搭梯子
		stairsNum += (height - currentHeight - 1) / dist
		currentHeight = height
	}
	return stairsNum
}

// @lc code=end

