/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	visited := map[int]bool{}
	for n != 1 && !visited[n] { // 当数字重复出现时即进入循环，退出
		visited[n] = true
		n = step(n)
	}
	return n == 1
}

func step(n int) (res int) {
	for n > 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return
}

// @lc code=end

