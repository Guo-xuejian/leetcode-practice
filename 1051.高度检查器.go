/*
 * @lc app=leetcode.cn id=1051 lang=golang
 *
 * [1051] 高度检查器
 */

// @lc code=start
package main

import "sort"

func heightChecker(heights []int) (res int) {
	expected := append([]int(nil), heights...)
	sort.Ints(expected)
	for i, v := range expected {
		if v != heights[i] {
			res++
		}
	}
	return
}

// @lc code=end
