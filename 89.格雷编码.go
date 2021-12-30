/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start
func grayCode(n int) []int {
	res := []int{0}
	head := 1
	for i := 0; i < n; i++ {
		// 二进制之间的相加，逐位递增，只会像相差一位
		// 001 ===> 011 ===> 111
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, head+res[j])
		}
		head <<= 1
	}
	return res
}

// @lc code=end

