/*
 * @lc app=leetcode.cn id=969 lang=golang
 *
 * [969] 煎饼排序
 */

// @lc code=start
func pancakeSort(arr []int) (res []int) {
	// 每次移动一个最大值至当前尾部
	for n := len(arr); n > 1; n-- {
		index := 0
		for i := 0; i < n; i++ { // 找当前最大值
			if arr[i] > arr[index] {
				index = i
			}
		}
		if index == n-1 { // 已经是最后一个不需要移动
			continue
		}
		m := index
		// +1 向上取整
		for i := 0; i < (m+1)/2; i++ { // 第一次翻转将当前最大值移动至首位
			arr[i], arr[m-i] = arr[m-i], arr[i]
		}
		for i := 0; i < n/2; i++ { // 第二次翻转将当前最大值移动至尾部
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
		res = append(res, index+1, n)
	}
	return
}

// @lc code=end

