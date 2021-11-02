/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	// 单引号和双引号不一样，在使用 byte 类型时，单引号才行
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 切片模拟栈
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if pairs[s[i]] > 0 {
			// 如果栈为空或者最后一个元素不是该右括号对应的左括号，那么明显不符合条件
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 符合条件，栈中出栈对应的配对左括号
			stack = stack[:len(stack)-1]
		} else {
			// 左括号入栈
			stack = append(stack, s[i])
		}
	}
	// 栈为空则全部匹配，栈中仍有元素意味着不符合条件
	return len(stack) == 0
}

// @lc code=end

