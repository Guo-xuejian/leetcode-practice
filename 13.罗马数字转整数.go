/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
var symbolsValuesMap = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	strLength := len(s)
	// 初始罗马数值 num, 临时数值记录上一个数字 lastNumber,返回值 total
	num, lastNumber, total := 0, 0, 0
	// 倒序遍历,从最小值开始往上加,遇到组合的那种也很好处理,减掉较小值就完成了实际的加总,比如 CM 900,从后往前先加上 1000,遇到 C 之后因为它比 M 小,total - 100 即可
	for i := strLength - 1; i >= 0; i-- {
		// 记住!!!直接获取字符串索引处值拿到的是 byte 值,需要处理
		char := fmt.Sprintf("%c", s[i])
		num = symbolsValuesMap[char]
		// 遇到了组合数字,需要在 toatal 的基础上减去这个值
		// 同时题目中给出的条件只会出现一个数字为两个字符的组合
		// 那么需要一个 lastNumber 记录
		if num < lastNumber {
			total -= num
		} else { // 剩下的就是正常的相加
			total += num
		}
		// 记录这一次遍历的值,实际上不需要担心 I,V,X等对M的影响,量级有差异,不会出现
		lastNumber = num
	}
	return total
}

// @lc code=end

