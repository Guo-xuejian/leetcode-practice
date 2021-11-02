/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) (res []int) {
	// 先排序方便处理
	sort.Ints(nums)
	numsLength := len(nums)
	temp := nums[0]
	count := 1
	for i := 1; i < numsLength; i++ {
		if nums[i] == temp { // 相等则计数加一
			count++
			continue
		}
		// 不相等判断 count
		if count < 2 { // 小于2，则出现一次加入结果
			res = append(res, temp)
		}
		// 覆盖 temp 和 count
		temp = nums[i]
		count = 1
	}
	if count < 2 { // 最后一个元素没有进行判定，符合要求则加入
		res = append(res, temp)
	}
	return
}

// @lc code=end

