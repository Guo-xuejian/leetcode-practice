/*
 * @lc app=leetcode.cn id=1550 lang=golang
 *
 * [1550] 存在连续三个奇数的数组
 */

// @lc code=start
func threeConsecutiveOdds(arr []int) bool {
	threeLimited := 0
	for _, v := range arr {
		if isOdd(v) {
			threeLimited++
		} else {
			threeLimited = 0
		}
		if threeLimited == 3 { // 满足即退出
			return true
		}
	}
	return false
}

func isOdd(num int) bool { // 奇偶判定
	if num%2 == 1 {
		return true
	}
	return false
}

// @lc code=end

