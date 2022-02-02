/*
 * @lc app=leetcode.cn id=335 lang=golang
 *
 * [335] 路径交叉
 */

// @lc code=start
func isSelfCrossing(distance []int) bool {
	n := len(distance)

	for i := 3; i < n; i++ {
		// 第一类相交
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}
		
		// 第二类相交
		if i == 4 && distance[3] == distance[1] && distance[4] >= distance[2] - distance[0] {
			return true
		}

		// 第三类相交
		if i >= 5 && distance[i - 3] - distance[i - 5] <= distance[i-1] && distance[i-1] <= distance[i-3] && distance[i] >= distance[i-2] - distance[i-4] && distance[i-2] > distance[i-4] {
			return true
		} 
	}
	// 没有相交
	return false
}
// @lc code=end

