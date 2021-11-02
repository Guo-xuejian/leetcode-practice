/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	// 初始化边界值，下一个字符是否允许为+-符号，下一个字符是否允许为空格
	maxInt, signAllowed, whitespaceAllowed := int64(2<<30), true, true
	// 初始化正负数判定符号,因为正数与负数的边界值有差异[−2**31,  2**31 − 1],同时初始化一个 digits 用来存储数字
	sign, digits := 1, []int{}
	for _, c := range s {
		// 是空格并且上一步判定允许这一次遍历出现空格===>结束此次遍历
		if c == ' ' && whitespaceAllowed {
			continue
		}
		if signAllowed { // 判定下一次遍历是否允许出现符号
			if c == '+' { // 出现 +
				signAllowed = false       // 下一个不能是符号
				whitespaceAllowed = false // 下一个不能是空格
				continue
			} else if c == '-' { // 出现 -
				sign = -1                 // 判定为负数
				signAllowed = false       // 下一个不能是符号
				whitespaceAllowed = false // 下一个不能是空格
				continue
			} // 不需要 else,因为只判定数字,而非算式,实际上这部分代码块只会运行一次
		}
		// 字符不在[0...9]内,说明未遇到数字或者数据字符串结束,结束  for 循环
		if c < '0' || c > '9' {
			break
		}
		// 1.允许的空格出现或无空格
		// 2.未出现+-或已经判定
		// 3.c 处在[0...9]范围内
		whitespaceAllowed, signAllowed = false, false // 正常数字出现了,那么下一个字符不能是符号或者空格
		digits = append(digits, int(c-48))            // 写入这次的 数字
	}
	// 初始化place为进制为个位代表的乘数1,(十位也就是10====>方便理解),int8 很有可能不足以支撑
	// 初始化 num = 0为最终返回数字的初始值
	var num, place int64
	place, num = 1, 0
	// 如果 0 出现在开头,那么需要丢弃
	// lastLeading0Index 记录最后一次 0 出现的索引,方便后续截断操作
	lastLeading0Index := -1
	for i, d := range digits { // 遍历上方获得的数字切片
		if d == 0 { // 如果数字为 0,那么记录索引
			lastLeading0Index = i
		} else { // 只要第一个数字不是 0 ,那么就是一个正常的数字,结束遍历
			break
		}
	}
	// 如果 0 开头,那么截断切片至切片第一个元素不是 0
	if lastLeading0Index > -1 {
		digits = digits[lastLeading0Index+1:]
	}
	// 同上面所讲,正数和负数的大小判定值存在差异,通过 sign 记录正负状态
	var rtnMax int64
	if sign > 0 { // 正数 ===> 2***31 - 1
		rtnMax = maxInt - 1
	} else { // 负数就是一开始初始化的值
		rtnMax = maxInt
	}
	digitsCount := len(digits)
	// 倒序遍历,倒序的好处就是进制的处理比较方便,从个位开始,逐位向上,也比较好记录 place,
	// 正向也可以,通过 digitsCount 确定进位情况合理确定 乘数place就可以
	for i := digitsCount - 1; i >= 0; i-- {
		num += int64(digits[i]) * place // num 记录真实数值
		place *= 10                     // place 进位向上变化
		if digitsCount-i > 10 || num > rtnMax {
			// 超过 32位有符号整数,截断为 int(int64(sign) * rtnMax)
			return int(int64(sign) * rtnMax)
		}
	}
	// 上面为方便处理都是正数,现在还原为真是正负状态
	num *= int64(sign)
	// 题中要求返回值类型 int
	return int(num)
}

// @lc code=end

