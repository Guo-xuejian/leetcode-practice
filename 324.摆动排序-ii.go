/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 */

// @lc code=start
func wiggleSort(nums []int) {
	n := len(nums)
	arr := append([]int{}, nums...)
	sort.Ints(arr)
	x := (n + 1) / 2
	for i, j, k := 0, x-1, n-1; i < n; i += 2 {
		nums[i] = arr[j]
		if i+1 < n {
			nums[i+1] = arr[k]
		}
		j--
		k--
	}
}

// @lc code=end

