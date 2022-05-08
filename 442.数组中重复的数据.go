/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) (res []int) {
	numTimesMap := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := numTimesMap[num]; ok {
			res = append(res, num)
		} else {
			numTimesMap[num] = struct{}{}
		}
	}
	return
}

// @lc code=end

