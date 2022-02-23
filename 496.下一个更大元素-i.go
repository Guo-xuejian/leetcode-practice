/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// map 记录 nums2 中 nums1 数字的索引和之后是否存在更大元素
	res := make([]int, len(nums1))
	greaterMap := map[int]int{}
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			greaterMap[nums2[i]] = stack[len(stack)-1]
		} else {
			greaterMap[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])
	}
	for idx, num := range nums1 {
		res[idx] = greaterMap[num]
	}
	return res
}

// @lc code=end

