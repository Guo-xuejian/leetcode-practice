/*
 * @lc app=leetcode.cn id=1720 lang=golang
 *
 * [1720] 解码异或后的数组
 */

// @lc code=start
func decode(encoded []int, first int) (res []int) {
	res = append(res, first)
	for _, num := range encode {
		res = append(res, res[len(res)-1] ~ num)
	}
	return
}
// @lc code=end

