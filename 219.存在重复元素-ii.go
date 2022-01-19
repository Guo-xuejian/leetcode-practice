/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	allMap := map[int]int{}
	for idx, val := range nums {
		if preIndex, ok := allMap[val]; ok {
			if idx-preIndex <= k {
				return true
			}
		}
		// 即使重复也希望保留最近的数字，因为要求小于，所以离得近值自然越小
		allMap[val] = idx
	}
	// 不满足条件
	return false
}

// @lc code=end

