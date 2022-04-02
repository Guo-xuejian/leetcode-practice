/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) (res int) {
	left, right := 0, n
	res = n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return
}

// @lc code=end

