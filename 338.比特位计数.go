/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(n int) (bits []int) {
	for i := 0; i < n+1; i++ {
		bits = append(bits, countOnes(i))
	}
	return
}

func countOnes(x int) (ones int) {
	for x > 0 {
		x &= x - 1 // 做位与，从高位出发，直到没有 1
		ones++
	}
	return
}

// @lc code=end

