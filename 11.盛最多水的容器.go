/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	// 初始化最大值 左右边界
	max, start, end := 0, 0, len(height)-1
	// 对撞指针不对撞，对撞之后时无意义的重复
	for start < end {
		// 获得宽度
		width := end - start
		// 盛水多少取决于哪边比较矮
		// high 会成为左右边界中较小的一个
		high := 0
		if height[start] < height[end] {
			// 左边界较矮
			high = height[start]
			start++
		} else {
			// 右边界较矮
			high = height[end]
			end--
		}

		// 计算临时面积并和 max 比较（超越则覆盖）
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}

// @lc code=end

