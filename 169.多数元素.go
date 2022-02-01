/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) (res int) {
	vote := 0
	for _, num := range nums {
		if vote == 0 {
			res = num
		}
		if res == num {
			vote++
		} else {
			vote--
		}
	}
	return
}
// @lc code=end

