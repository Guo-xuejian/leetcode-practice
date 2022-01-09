/*
 * @lc app=leetcode.cn id=1437 lang=golang
 *
 * [1437] 是否所有 1 都至少相隔 k 个元素
 */

// @lc code=start
func kLengthApart(nums []int, k int) bool {
	if k == 0 { // 间隔为 0 天生满足,
		return true
	}
	if len(nums) < k { // 长度小于间隔自然不满足
		return false
	}
	preIndex := 0
	existStatus := false
	for idx, val := range nums {
		if val == 1 {
			if !existStatus {
				existStatus = true
				preIndex = idx
				continue
			}
			if idx-preIndex-1 >= k {
				preIndex = idx
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end

