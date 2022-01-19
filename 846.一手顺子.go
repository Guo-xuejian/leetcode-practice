/*
 * @lc app=leetcode.cn id=846 lang=golang
 *
 * [846] 一手顺子
 */

// @lc code=start
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize > 0 {
		return false
	}
	numMap := map[int]int{}
	for _, val := range hand {
		numMap[val]++
	}
	sort.Ints(hand)
	for _, val := range hand {
		if numMap[val] == 0 { // 已经被抵消
			continue
		}
		// 作为起点查看是否满足,其实也就是相同的部分会另起一个牌堆，匹配完成则成功
		for i := val; i < val+groupSize; i++ {
			if numMap[i] == 0 {
				return false
			}
			numMap[i]-- // 抵消
		}
	}
	// 完全匹配成功
	return true
}

// @lc code=end

