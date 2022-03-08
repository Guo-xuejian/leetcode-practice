/*
 * @lc app=leetcode.cn id=798 lang=golang
 *
 * [798] 得分最高的最小轮调
 */

// @lc code=start
func bestRotation(nums []int) int {
	n := len(nums)
	diffs := make([]int, n)
	for i, num := range nums {
		low := (i + 1) % n
		high := (i - num + n + 1) % n
		diffs[low]++
		diffs[high]--
		if low >= high {
			diffs[0]++
		}
	}
	score, maxScore, idx := 0, 0, 0
	for i, diff := range diffs {
		score += diff
		if score > maxScore {
			maxScore, idx = score, i
		}
	}
	return idx
}

// @lc code=end

