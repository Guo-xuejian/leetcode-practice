/*
 * @lc app=leetcode.cn id=1470 lang=golang
 *
 * [1470] 重新排列数组
 */

// @lc code=start
func shuffle(nums []int, n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, nums[i])
		res = append(res, nums[i+n])
	}
	return
}

// 2022-08-29
func shuffle(nums []int, n int) (res []int) {
	res = make([]int, 2*n)
	// half is enough
	for idx := range nums[:n] {
		res[2*idx] = nums[idx]
		res[2*idx+1] = nums[n+idx]
	}
	return
}

// @lc code=end

