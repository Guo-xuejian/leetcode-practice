// 剑指 Offer 10- II. 青蛙跳台阶问题
// 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

// 示例 1：

// 输入：n = 2
// 输出：2
// 示例 2：

// 输入：n = 7
// 输出：21
// 示例 3：

// 输入：n = 0
// 输出：1
// 提示：

// 0 <= n <= 100

func numWays(n int) int {
	// 递归必超时
	// if n == 0 {
	//     return 1
	// }else if n <= 2 {
	//     return n
	// }
	// return numWays(n - 1) + numWays(n - 2)
	if n == 0 || n == 1 {
		return 1
	}
	prePre, pre := 1, 1
	for n > 1 {
		prePre, pre = pre, (pre+prePre)%1000000007
		n--
	}
	return pre % 1000000007
}