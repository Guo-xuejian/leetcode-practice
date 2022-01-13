/*
 * @lc app=leetcode.cn id=910 lang=golang
 *
 * [910] 最小差值 II
 */

// @lc code=start
func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	length := len(nums)
	res := nums[length-1] - nums[0]
	p1 := nums[0] + k
	p4 := nums[length-1] - k
	for i := 0; i < length-1; i++ {
		p2 := nums[i] + k
		p3 := nums[i+1] - k
		h := max(p2, p4)
		l := min(p1, p3)
		if res > h-l {
			res = h - l
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

