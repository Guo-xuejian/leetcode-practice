/*
 * @lc app=leetcode.cn id=1414 lang=golang
 *
 * [1414] 和为 K 的最少斐波那契数字数目
 */

// @lc code=start
func findMinFibonacciNumbers(k int) (res int) {
	fibSlice := []int{1, 1}
	for fibSlice[len(fibSlice)-1] < k {
		fibSlice = append(fibSlice, fibSlice[len(fibSlice)-1]+fibSlice[len(fibSlice)-2])
	}

	// 反向查找，优先找较大的数字来替代，最大、次大、较大依次填充至等于 k，存在 1 所以必定可以凑齐
	index := len(fibSlice) - 1
	for k > 0 {
		if fibSlice[index] <= k { // 找到较大的
			k -= fibSlice[index]
			res++
		}
		index-- // 指针前移寻找下一个较大的数字
	}
	return
}

// @lc code=end

