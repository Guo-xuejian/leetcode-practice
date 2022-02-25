/*
 * @lc app=leetcode.cn id=1588 lang=golang
 *
 * [1588] 所有奇数长度子数组的和
 */

// @lc code=start
func sumOddLengthSubarrays(arr []int) (res int) {
	n := len(arr)
	for start := 0; start < n; start++ {
		length := 1
		for start+length < n+1 {
			for _, num := range arr[start : start+length] {
				res += num
			}
			length += 2
		}
	}
	return
}

// @lc code=end

