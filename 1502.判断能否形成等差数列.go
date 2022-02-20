/*
 * @lc app=leetcode.cn id=1502 lang=golang
 *
 * [1502] 判断能否形成等差数列
 */

// @lc code=start
func canMakeArithmeticProgression(arr []int) bool {
	// 排序后判定
	length := len(arr)
	if length == 2 {
		return true
	}
	sort.Ints(arr)
	lap := arr[1] - arr[0] // 等差值
	for i := 2; i < length; i++ {
		if arr[i]-arr[i-1] != lap {
			return false
		}
	}
	return true
}

// @lc code=end

