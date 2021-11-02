/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */

// @lc code=start
func constructRectangle(area int) []int {
	// 极限值（不要求整数值）会相等，
	// 那么 w 就只能往下减了，
	// 寻找到一个符合条件的数字，从相等出发
	// ，运算次数最少
	w := int(math.Sqrt(float64(area)))
	for area%w > 0 {
		w--
	}
	return []int{area / w, w}
}

// @lc code=end

