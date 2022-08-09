/*
 * @lc app=leetcode.cn id=640 lang=golang
 *
 * [640] 求解方程
 */

// @lc code=start
func solveEquation(equation string) string {
	lineSplit := strings.Split(equation, "=")
	lineLeft, lineRight := lineSplit[0], lineSplit[1]
	xLeft, numLeft := splitLine(lineLeft)
	xRight, numRight := splitLine(lineRight)

	x := xLeft - xRight
	num := numRight - numLeft
	if x == 0 {
		if num == 0 {
			return "Infinite solutions"
		} else {
			return "No Solution"
		}
	}
	return "x=" + strconv.Itoa(num/x)
}

func splitLine(line string) (x, num int) {
	if line[0] == 'x' {
		line = "1" + line
	}
	text := strings.Replace(strings.Replace(strings.Replace(line, "+x", "+1x", -1), "-x", "-1x", -1), "-", "+-", -1)
	textSlice := strings.Split(text, "+")
	for _, one := range textSlice {
		if len(one) > 0 {
			if one[len(one)-1] == 'x' {
				xNow, _ := strconv.Atoi(string(one[:len(one)-1]))
				x += xNow
			} else {
				numNow, _ := strconv.Atoi(string(one))
				num += numNow
			}
		}
	}
	return

}

// 2022-08-10
func solveEquation(equation string) (res string) {
	equation = strings.Replace(equation, "-", "+-", -1)
	allSlice := strings.Split(equation, "=")
	left, right := strings.Split(allSlice[0], "+"), strings.Split(allSlice[1], "+")
	varNum, constNum := 0, 0
	currLength := 0
	for _, one := range left {
		if one == "" {
			continue
		}
		currLength = len(one)
		if one[currLength-1] == 'x' {
			varNum += getNum(one[:currLength-1])
		} else {
			constNum -= getNum(one)
		}
	}
	for _, one := range right {
		if one == "" {
			continue
		}
		currLength = len(one)
		if one[currLength-1] == 'x' {
			varNum -= getNum(one[:currLength-1])
		} else {
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
	if numString == "" {
		return 1
	} else if numString == "-" {
		return -1
	}
	multi := 1
	for i := len(numString) - 1; i >= 0; i-- {
		if numString[i] == '+' {
			continue
		} else if numString[i] == '-' {
			res *= -1
		} else {
			res += int(byte(numString[i])-'0') * multi
			multi *= 10
		}
	}
	return
}

// @lc code=end

