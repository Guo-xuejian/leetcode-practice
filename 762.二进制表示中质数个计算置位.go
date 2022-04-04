/*
 * @lc app=leetcode.cn id=762 lang=golang
 *
 * [762] 二进制表示中质数个计算置位
 */

// @lc code=start
func countPrimeSetBits(left int, right int) (res int) {
	for num := left; num <= right; num++ {
		if isPrime(countBitsOne(num)) {
			res++
		}
	}
	return
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	i := 2
	for i*i <= num {
		if num%i == 0 {
			return false
		}
		i++
	}
	return true
}

func countBitsOne(num int) (res int) {
	for num != 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return
}

// @lc code=end

