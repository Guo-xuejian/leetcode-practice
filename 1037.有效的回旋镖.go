/*
 * @lc app=leetcode.cn id=1037 lang=golang
 *
 * [1037] 有效的回旋镖
 */

// @lc code=start
package main

func isBoomerang(points [][]int) bool {
    v1 := [2]int{points[1][0] - points[0][0], points[1][1] - points[0][1]}
    v2 := [2]int{points[2][0] - points[0][0], points[2][1] - points[0][1]}
    return v1[0]*v2[1]-v1[1]*v2[0] != 0
}
// @lc code=end

