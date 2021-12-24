// 给定非负整数数组 heights ，数组中的数字用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
// 输入：heights = [2,1,5,6,2,3]
// 输出：10
// 解释：最大的矩形为图中红色区域，面积为 10

func largestRectangleArea(heights []int) (maxArea int) {
	length := len(heights)
	for i := 0; i < length; i++ {
		// 以每一个直方作为起点，左右遍历获取单调栈，最终计算面积与 maxArea 比较取较大值
		currentHeight := heights[i]
		left := i - 1
		right := i + 1
		if length*currentHeight > maxArea {
			// 这个判定可以帮助之后的遍历减少计算，由于之前的计算已经包含
			// 或者说大于该直方面积，所以不需要再次进行额外计算，剪枝操作，有意思
			for left >= 0 && heights[left] >= currentHeight {
				left--
			}
			for right <= length-1 && heights[right] >= currentHeight {
				right++
			}
			maxArea = max(maxArea, (right-left-1)*currentHeight)
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}