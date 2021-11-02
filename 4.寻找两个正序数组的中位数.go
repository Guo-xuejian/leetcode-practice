/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	median1, median2 := 0, 0
	for left <= right {
		i := (left + right) / 2
		j := (m+n+1)/2 - i
		nums_im1 := math.MinInt32
		if i != 0 {
			nums_im1 = nums1[i-1]
		}
		nums_i := math.MaxInt32
		if i != m {
			nums_i = nums1[i]
		}
		nums_jm1 := math.MinInt32
		if j != 0 {
			nums_jm1 = nums2[j-1]
		}
		nums_j := math.MaxInt32
		if j != n {
			nums_j = nums2[j]
		}
		if nums_im1 <= nums_j {
			median1 = max(nums_im1, nums_jm1)
			median2 = min(nums_i, nums_j)
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (m+n)%2 == 0 {
		return float64(median1+median2) / 2.0
	}
	return float64(median1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

