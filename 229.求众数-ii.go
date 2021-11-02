/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 */

// @lc code=start
func majorityElement(nums []int) (res []int) {
	lengthOfNums := len(nums)
	threeLimitNum := int(lengthOfNums / 3)
	sort.Ints(nums) // 排序数组方便处理
	temp := nums[0]
	count := 1
	for i := 1; i < lengthOfNums; i++ { // 1 剔除了第一个元素多计数一次
		if nums[i] == temp { // 和之前的元素相等计数加一
			count++
			continue
		}
		if count > threeLimitNum { // 不相等且满足条件则添加进入结果
			res = append(res, temp)
		}
		// 满不满足条件都要在遇到不同元素时覆盖 temp 值，并将 count 归 1
		temp = nums[i]
		count = 1
	}
	if count > threeLimitNum { // 数组的最后几个都为同一元素会导致条件判定未发生，需要特殊处理
		res = append(res, temp)
	}
	return
}

// @lc code=end

