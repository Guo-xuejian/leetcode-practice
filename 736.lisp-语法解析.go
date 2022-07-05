/*
 * @lc app=leetcode.cn id=736 lang=golang
 *
 * [736] Lisp 语法解析
 */

// @lc code=start
func evaluate(expression string) int {
	i, n := 0, len(expression)
	parseVar := func() string {
		i0 := i
		for i < n && expression[i] != ' ' && expression[i] != ')' {
			i++
		}
		return expression[i0:i]
	}
	parseInt := func() int {
		sign, x := 1, 0
		if expression[i] == '-' {
			sign = -1
			i++
		}
		for i < n && unicode.IsDigit(rune(expression[i])) {
			x = x*10 + int(expression[i]-'0')
			i++
		}
		return sign * x
	}

	scope := map[string][]int{}
	var innerEvaluate func() int
	innerEvaluate = func() (ret int) {
		if expression[i] != '(' { // 非表达式，可能为：整数或变量
			if unicode.IsLower(rune(expression[i])) { // 变量
				vals := scope[parseVar()]
				return vals[len(vals)-1]
			}
			return parseInt() // 整数
		}
		i++                       // 移除左括号
		if expression[i] == 'l' { // "let" 表达式
			i += 4 // 移除 "let "
			vars := []string{}
			for {
				if !unicode.IsLower(rune(expression[i])) {
					ret = innerEvaluate() // let 表达式的最后一个 expr 表达式的值
					break
				}
				vr := parseVar()
				if expression[i] == ')' {
					vals := scope[vr]
					ret = vals[len(vals)-1] // let 表达式的最后一个 expr 表达式的值
					break
				}
				vars = append(vars, vr)
				i++ // 移除空格
				scope[vr] = append(scope[vr], innerEvaluate())
				i++ // 移除空格
			}
			for _, vr := range vars {
				scope[vr] = scope[vr][:len(scope[vr])-1] // 清除当前作用域的变量
			}
		} else if expression[i] == 'a' { // "add" 表达式
			i += 4 // 移除 "add "
			e1 := innerEvaluate()
			i++ // 移除空格
			e2 := innerEvaluate()
			ret = e1 + e2
		} else { // "mult" 表达式
			i += 5 // 移除 "mult "
			e1 := innerEvaluate()
			i++ // 移除空格
			e2 := innerEvaluate()
			ret = e1 * e2
		}
		i++ // 移除右括号
		return
	}
	return innerEvaluate()
}

// @lc code=end

