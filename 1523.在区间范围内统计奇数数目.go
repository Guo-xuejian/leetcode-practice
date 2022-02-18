/*
 * @lc app=leetcode.cn id=1523 lang=golang
 *
 * [1523] 在区间范围内统计奇数数目
 */

// @lc code=start
func countOdds(low int, high int) (res int) {
	// for low <= high {
	//     if high % 2 != 0 {
	//         res++
	//     }
	//     if low % 2 != 0 && low != high {
	//         res++
	//     }
	//     high--
	//     low++
	// }
	// return
	if high%2 == 1 {
		high++
	}
	return (high >> 1) - (low >> 1)
}

// @lc code=end

