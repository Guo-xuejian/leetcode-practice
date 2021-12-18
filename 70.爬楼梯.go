/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	// 1.递归
	// if n <= 2 {
	// 	return n
	// }
	// // 但是递归调用的时间时间开销和空间开销都很不乐观
	// return climbStairs(n-1) + climbStairs(n-2)

	// 2.数组模拟斐波那契数列
	// fibArray := []int{}
	// for _ := n+1 {
	// 	fibArray = append(fibArray, 0)
	// }
	// fibArray[0] = 1
	// fibArray[1] = 1
	// for i := 2; i < n + 1; i++ {
	// 	fibArray[i] = fibArray[i-1] + fibArray[i-2]
	// }
	// return fibArray[-1]	// 其实也存在较大空间开销

	// 3.究极变量交换
	pre := 1
	after := 1
	for i := 2; i < n+1; i++ {
		pre, after = after, after+pre
	}
	return after
}

// @lc code=end

