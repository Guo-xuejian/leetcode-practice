/*
 * @lc app=leetcode.cn id=1217 lang=golang
 *
 * [1217] 玩筹码
 */

// @lc code=start
func minCostToMoveChips(position []int) int {
	// 偶数位是等价的，奇数位也可以通过移动偶数位等价
	// 所以实际上偶数位上的点可以看做同一个点，奇数位同样，比较奇偶数那个小一些，移动的代价就最小
	oddNum := 0
	for _, pos := range position {
		oddNum += pos % 2
	}
	return min(oddNum, len(position)-oddNum)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}


// func minCostToMoveChips(position []int) int {
//     cnt := [2]int{}
//     for _, p := range position {
//         cnt[p%2]++
//     }
//     return min(cnt[0], cnt[1])
// }

// func min(a, b int) int {
//     if a > b {
//         return b
//     }
//     return a
// }

// @lc code=end

