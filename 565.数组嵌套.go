/*
 * @lc app=leetcode.cn id=565 lang=golang
 *
 * [565] 数组嵌套
 */

// @lc code=start
func arrayNesting(nums []int) (ans int) {
	vis := make([]bool, len(nums))
	for i := range vis {
		cnt := 0
		for !vis[i] {
			vis[i] = true
			i = nums[i]
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return
}

// @lc code=end

