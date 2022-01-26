// 面试题 05.04. 下一个数
// 下一个数。给定一个正整数，找出与其二进制表达式中1的个数相同且大小最接近的那两个数（一个略大，一个略小）。

// 示例1:

//  输入：num = 2（或者0b10）
//  输出：[4, 1] 或者（[0b100, 0b1]）
// 示例2:

//  输入：num = 1
//  输出：[2, -1]
// 提示:

// num的范围在[1, 2147483647]之间；
// 如果找不到前一个或者后一个满足条件的正数，那么输出 -1。

func findClosedNumbers(num int) []int {
	var maxNum, minNum int64
	binString := strconv.FormatInt(int64(num), 2)
	n := len(binString)
	backNum := binString[n-1]
	var ones, zeros int // 数字 0 1 数量
	if backNum == '0' {
		zeros = 1
	} else {
		ones = 1
	}
	for i := n - 2; i >= 0; i-- {
		curr := binString[i]

		if curr == '1' {
			ones++
			if minNum == 0 && i != 0 && backNum == '0' { // 10 交换得到较小值，尽量大，1 中间，0 后置

				temp := string(binString[:i]) + "01"
				for j := 0; j < ones-1; j++ {
					temp += "1"
				}
				for j := 0; j < zeros-1; j++ {
					temp += "0"
				}
				minNum, _ = strconv.ParseInt(temp, 2, 0)
			}
		} else {
			zeros++
			if maxNum == 0 && backNum == '1' { // 01 交换得到较大值，尽量小，0 中间，其余 1 后置
				temp := string(binString[:i]) + "10"
				for j := 0; j < zeros-1; j++ {
					temp += "0"
				}
				for j := 0; j < ones-1; j++ {
					temp += "1"
				}
				maxNum, _ = strconv.ParseInt(temp, 2, 0)
			}
		}
		backNum = curr                // 重置拖尾字符
		if maxNum > 0 && minNum > 0 { // 直接满足就返回
			return []int{int(maxNum), int(minNum)}
		}
	}
	if minNum == 0 { // 找不到可以删除一个，删除后尽量大，删除 0,同时 1 前置
		if zeros > 0 {
			temp := ""
			for j := 0; j < ones; j++ {
				temp += "1"
			}
			for j := 0; j < zeros-1; j++ {
				temp += "0"
			}
			minNum, _ = strconv.ParseInt(temp, 2, 0)
		} else { // 没有 0 就没法删
			minNum = -1
		}
	}

	if maxNum == 0 { // 添加一个，尽量小，添加 0,其余 1 后置
		temp := "1"
		for j := 0; j < zeros+1; j++ {
			temp += "0"
		}
		for j := 0; j < ones-1; j++ {
			temp += "1"
		}
		maxNum, _ = strconv.ParseInt(temp, 2, 0)
		if maxNum > 2147483647 { // 有范围限制
			maxNum = -1
		}
	}
	return []int{int(maxNum), int(minNum)}
}