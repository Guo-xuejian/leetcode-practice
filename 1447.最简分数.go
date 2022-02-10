/*
 * @lc app=leetcode.cn id=1447 lang=golang
 *
 * [1447] 最简分数
 */

// @lc code=start
func simplifiedFractions(n int) (res []string) {
	for down := 2; down < n + 1; down++ {
		for up := 1; up < down; up++ {
			if gcd(up, down) == 1 {
				res = append(res, strconv.Itoa(up) + "/" + strconv.Itoa(down))
			}
		}
	}
	return
}


func gcd(x, y int) int {
	for x != 0 {
		x, y = y % x, x
	}
	return y
}
// @lc code=end

