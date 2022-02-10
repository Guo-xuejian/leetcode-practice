/*
 * @lc app=leetcode.cn id=1984 lang=golang
 *
 * [1984] 学生分数的最小差值
 */

// @lc code=start
func minimumDifference(nums []int, k int) int {
	if k == 1 { // 只取一个，那就是 num - num = 0
		return 0
	}
	sort.Ints(nums) // 排序方便处理
	length := len(nums)
	res := nums[length-1] // res 取小，所以给个最大值
	left := 0
	for left <= len(nums)-k {
		// 最大值和最小值差值与 res 比较
		res = min(res, nums[left+k-1]-nums[left])
		left++
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

