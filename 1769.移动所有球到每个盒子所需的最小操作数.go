/*
 * @lc app=leetcode.cn id=1769 lang=golang
 *
 * [1769] 移动所有球到每个盒子所需的最小操作数
 */

// @lc code=start
func minOperations(boxes string) (res []int) {
	length := len(boxes)

	var findBoxes = func(index int) (res int) {
		left, right := index - 1, index + 1

		for left >= 0 || right <= length - 1 {
			if left >= 0 {
				if boxes[left] == '1' {
					res += index - left
				}
				left--
			}
			if right <= length - 1 {
				if boxes[right] == '1' {
					res += right - index
				}
				right++
			}
		}
		return
	}

	for i := 0; i < length; i++ {
		res = append(res, findBoxes(i))
	}
	return
}
// @lc code=end

