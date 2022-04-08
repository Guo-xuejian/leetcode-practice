/*
 * @lc app=leetcode.cn id=1855 lang=golang
 *
 * [1855] 下标对中的最大距离
 */

// @lc code=start
func maxDistance(nums1 []int, nums2 []int) (res int) {
	m, n := len(nums1), len(nums2)
	index := 0
	for i := 0; i < n; i++ {
		for index < m && nums1[index] > nums2[i] {
			index++
		}
		if index < m {
			res = max(res, i-index)
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

