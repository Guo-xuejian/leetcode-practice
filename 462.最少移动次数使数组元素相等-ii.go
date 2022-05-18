/*
 * @lc app=leetcode.cn id=462 lang=golang
 *
 * [462] 最少移动次数使数组元素相等 II
 */

// @lc code=start
func minMoves2(nums []int) (res int) {
	sort.Ints(nums)
	middleNum := nums[len(nums)/2]
	for _, num := range nums {
		res += abs(num - middleNum)
	}
	return
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

// @lc code=end

