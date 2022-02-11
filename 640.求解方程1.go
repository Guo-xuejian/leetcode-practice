/*
 * @lc app=leetcode.cn id=640 lang=golang
 *
 * [640] 求解方程
 */

// @lc code=start
func solveEquation(equation string) (res string) {
	// 将所有的 - 替换为 +-，后续方便以 + 切分
	equation = strings.Replace(equation, "-", "+-", -1)
	allSlice := strings.Split(equation, "=")
	// 根据 + 切分
	left, right := strings.Split(allSlice[0], "+"), strings.Split(allSlice[1], "+")
	varNum, constNum := 0, 0 // 变量系数和常数
	currLength := 0
	for _, one := range left {
		if one == "" {
			continue
		}
		currLength = len(one)
		if one[currLength-1] == 'x' {
			// 等式左边变量系数为正
			varNum += getNum(one[:currLength-1])
		} else {
			// 等式左边常数为负
			constNum -= getNum(one)
		}
	}
	for _, one := range right {
		if one == "" {
			continue
		}
		currLength = len(one)
		if one[currLength-1] == 'x' {
			// 等式右边变量系数为负
			varNum -= getNum(one[:currLength-1])
		} else {
			// 等式右边常数为正
			constNum += getNum(one)
		}
	}
	if varNum == 0 && constNum == 0 {
		res = "Infinite solutions"
	} else if varNum == 0 {
		res = "No solution"
	} else {
		res = "x=" + strconv.Itoa(constNum/varNum)
	}
	return
}

func getNum(numString string) (res int) {
	if numString == "" { // x 情况
		return 1
	} else if numString == "-" { // -x 情况
		return -1
	}
	multi := 1 // 倍数
	for i := len(numString) - 1; i >= 0; i-- {
		if numString[i] == '+' { // + 不需要处理
			continue
		} else if numString[i] == '-' { // - 变为负数
			res *= -1
		} else {
			res += int(byte(numString[i])-'0') * multi
			multi *= 10 // 倍数递增
		}
	}
	return
}

// @lc code=end

