/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	costLength := len(cost)
	// pre 代表前两位，curr 代表前一个
	pre, curr := 0, 0
	for i := 2; i < costLength+1; i++ {
		// 每一次都取代价最小值，动态规划，对已知条件的遍历，取出最合适的方案
		pre, curr = curr, min(curr+cost[i-1], pre+cost[i-2])
	}
	return curr
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

