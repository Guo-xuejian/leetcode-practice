// 剑指 Offer 44. 数字序列中某一位的数字
// 数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

// 请写一个函数，求任意第n位对应的数字。

// 示例 1：

// 输入：n = 3
// 输出：3
// 示例 2：

// 输入：n = 11
// 输出：0

// 限制：

// 0 <= n < 2^31

func findNthDigit(n int) int {
	digit, digitNum, count := 1, 1, 9
	for n > count {
		n -= count
		digit++
		digitNum *= 10
		count = 9 * digit * digitNum
	}
	num := digitNum + (n-1)/digit

	index := (n - 1) % digit

	numStr := strconv.Itoa(num)
	return int(numStr[index] - '0')
}
