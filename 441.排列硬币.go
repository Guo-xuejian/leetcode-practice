/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	if n <= 1 { // 快速判断并退出
		return n
	}
	// 定义 res 返回值
	res := 0
	totalNum := n // totalNum 记录 还剩下多少个硬币
	// 由于每一行的硬币数量与行号有关
	// 与 for 循环的索引值相似，通过 i+1 来模拟每个应该存在的行所必须需要的硬币数量
	for i := 0; i < n; i++ {
		totalNum -= i + 1 // 遍历一减去对应的数量
		if totalNum >= 0 {
			continue
		} else { // // 当相减至当前所剩下的硬币数量不足以支撑这一行对应的硬币数量，记录索引，退出
			res = i
			break
		}
	}
	return res
}

// @lc code=end

