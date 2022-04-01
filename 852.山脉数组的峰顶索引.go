/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	m := len(arr)
	left, right := 0, m-1
	for left <= right {
		mid := (right-left)>>1 + left
		// fmt.Println(left, right, mid)
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] > arr[mid-1] {
			left = mid
		} else {
			right = mid
		}
	}
	return -1
}

// @lc code=end

