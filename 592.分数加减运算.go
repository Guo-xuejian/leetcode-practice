/*
 * @lc app=leetcode.cn id=592 lang=golang
 *
 * [592] 分数加减运算
 *
 * https://leetcode.cn/problems/fraction-addition-and-subtraction/description/
 *
 * algorithms
 * Medium (52.82%)
 * Likes:    121
 * Dislikes: 0
 * Total Accepted:    22.1K
 * Total Submissions: 36.6K
 * Testcase Example:  '"-1/2+1/2"'
 *
 * 给定一个表示分数加减运算的字符串 expression ，你需要返回一个字符串形式的计算结果。
 *
 * 这个结果应该是不可约分的分数，即最简分数。 如果最终结果是一个整数，例如 2，你需要将它转换成分数形式，其分母为 1。所以在上述例子中, 2
 * 应该被转换为 2/1。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: expression = "-1/2+1/2"
 * 输出: "0/1"
 *
 *
 * 示例 2:
 *
 *
 * 输入: expression = "-1/2+1/2+1/3"
 * 输出: "1/3"
 *
 *
 * 示例 3:
 *
 *
 * 输入: expression = "1/3-1/2"
 * 输出: "-1/6"
 *
 *
 *
 *
 * 提示:
 *
 *
 * 输入和输出字符串只包含 '0' 到 '9' 的数字，以及 '/', '+' 和 '-'。
 * 输入和输出分数格式均为 ±分子/分母。如果输入的第一个分数或者输出的分数是正数，则 '+' 会被省略掉。
 * 输入只包含合法的最简分数，每个分数的分子与分母的范围是  [1,10]。 如果分母是1，意味着这个分数实际上是一个整数。
 * 输入的分数个数范围是 [1,10]。
 * 最终结果的分子与分母保证是 32 位整数范围内的有效整数。
 *
 *
 */

// @lc code=start
func fractionAddition(expression string) string {
	denominator, numerator := 0, 1 // 分子，分母
	for i, n := 0, len(expression); i < n; {
		// 读取分子
		denominator1, sign := 0, 1
		if expression[i] == '-' || expression[i] == '+' {
			if expression[i] == '-' {
				sign = -1
			}
			i++
		}
		for i < n && unicode.IsDigit(rune(expression[i])) {
			denominator1 = denominator1*10 + int(expression[i]-'0')
			i++
		}
		denominator1 = sign * denominator1
		i++

		// 读取分母
		numerator1 := 0
		for i < n && unicode.IsDigit(rune(expression[i])) {
			numerator1 = numerator1*10 + int(expression[i]-'0')
			i++
		}

		denominator = denominator*numerator1 + denominator1*numerator
		numerator *= numerator1
	}
	if denominator == 0 {
		return "0/1"
	}
	g := gcd(abs(denominator), numerator)
	return fmt.Sprintf("%d/%d", denominator/g, numerator/g)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

