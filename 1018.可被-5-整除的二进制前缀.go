/*
 * @lc app=leetcode.cn id=1018 lang=golang
 *
 * [1018] 可被 5 整除的二进制前缀
 */

// @lc code=start
package main

func prefixesDivBy5(nums []int) []bool {
	res := make([]bool, len(nums))
	preNum := 0
	for i, v := range nums {
		// 只保留 0 或者不被 5 整除的部分，避免超出 int 的限制大小
		preNum = (preNum<<1 + v) % 5
		if preNum == 0 {
			res[i] = true
		}
	}
	return res
}

// @lc code=end

