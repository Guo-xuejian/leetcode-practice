/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 从后往前写入，较大先写入，较小后写入
	i, j, z := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 { // 循环结束时可能存在其中一个未能完全写入
		if nums1[i] < nums2[j] {
			nums1[z] = nums2[j]
			j--
		} else {
			nums1[z] = nums1[i]
			i--
		}
		z--
	}
	// 存在这样一种情况，nums1 全部先写完或者写完时 num2 未完全写入，因为循环会在其中一个索引越界时退出
	for i := 0; i < j+1; i++ {
		nums1[i] = nums2[i]
	}
}

// @lc code=end

