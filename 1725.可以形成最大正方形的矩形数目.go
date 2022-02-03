/*
 * @lc app=leetcode.cn id=1725 lang=golang
 *
 * [1725] 可以形成最大正方形的矩形数目
 */

// @lc code=start
func countGoodRectangles(rectangles [][]int) (res int) {
	maxLength := 0
	for i := 0; i < len(rectangles); i++ {
		length, width := rectangles[i][0], rectangles[i][1]
		// 长方形的短边即为最大切分正方形边长
		num := min(length, width)
		if num == maxLength {
			res++
		} else if num > maxLength {	// 更大值出现时重置 res 和 maxLength
			res = 1
			maxLength = num
		}
	}
	return
}


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

