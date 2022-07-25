/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
// @lc code=start
package main

func twoSum(nums []int, target int) []int {
	length := len(nums)
	res := []int{}
label:
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i != j {
				if nums[i]+nums[j] == target {
					res = append(res, i, j)
					break label
				}
			}
		}
	}
	return res
}

// @lc code=end

