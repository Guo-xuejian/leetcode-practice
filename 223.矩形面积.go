/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 */

// @lc code=start
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	// 先计算出两个矩形的面积
	area1 := (ax2 - ax1) * (ay2 - ay1)
	area2 := (bx2 - bx1) * (by2 - by1)

	// 获取长宽,右顶点两值都大于左顶点
	overlapWidth := min(ax2, bx2) - max(ax1, bx1)
	overlapHeight := min(ay2, by2) - max(ay1, by1)
	// 如果出现负数则意味着两个矩形并没有重合区域.最终计算结果应该为 0
	overlapArea := max(overlapWidth, 0) * max(overlapHeight, 0)
	// 应该是减去，而不是加上
	return area1 + area2 - overlapArea
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

