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

// @lc code=end

