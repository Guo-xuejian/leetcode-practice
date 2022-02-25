/*
 * @lc app=leetcode.cn id=1232 lang=golang
 *
 * [1232] 缀点成线
 */

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	// 执行用时：4 ms, 在所有 Go 提交中击败了93.75%的用户
	// 内存消耗：3.4 MB, 在所有 Go 提交中击败了79.69%的用户
	// 满足同线方程的两个点： x1*y2 - x2*y1 = 0
	deltaX, deltaY := coordinates[0][0], coordinates[0][1]
	// 做移动使得第一个点过坐标原点，后续坐标同等操作
	x1, y1 := coordinates[1][0]-deltaX, coordinates[1][1]-deltaY
	for _, point := range coordinates[2:] { // 两点确定一条直线，从第三个元素开始判定
		x2, y2 := point[0]-deltaX, point[1]-deltaY // 同理移动
		if x1*y2-x2*y1 != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

