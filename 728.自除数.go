/*
 * @lc app=leetcode.cn id=728 lang=golang
 *
 * [728] 自除数
 */

// @lc code=start
func selfDividingNumbers(left int, right int) (res []int) {
	for num := left; num < right + 1; num++ {
		if valid(num) {
			res = append(res, num)
		}
	}
	return
}


func valid(number int) bool {
	for _, curr := range strconv.Itoa(number) {
		currInt := int(curr) - int('0')
		if currInt == 0 || number % currInt != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

