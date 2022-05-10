// 剑指 Offer 17. 打印从1到最大的n位数
// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

// 示例 1:

// 输入: n = 1
// 输出: [1,2,3,4,5,6,7,8,9]

// 说明：

// 用返回一个整数列表来代替打印
// n 为正整数

func printNumbers(n int) []int {
	res := make([]int, int(math.Pow10(n)-1))
	count := 0
	var dfs func(index int, num []byte, digit int)
	dfs = func(index int, num []byte, digit int) {
		if index == digit {
			tmp, _ := strconv.Atoi(string(num))
			res[count] = tmp
			count++
			return
		}
		for i := '0'; i <= '9'; i++ {
			num[index] = byte(i)
			dfs(index+1, num, digit)
		}
	}
	for digit := 1; digit <= n; digit++ {
		for first := '1'; first <= '9'; first++ {
			num := make([]byte, digit)
			num[0] = byte(first)
			dfs(1, num, digit)
		}
	}
	return res
}
