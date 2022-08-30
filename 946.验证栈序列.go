/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))
	popIndex := 0
	for _, num := range pushed {
		// get in stack
		stack = append(stack, num)
		// check if current situation could work,
		// yes one element out stack,
		// no continue(the pop one may be the next one in pushed)
		for len(stack) != 0 && stack[len(stack)-1] == popped[popIndex] {
			stack = stack[:len(stack)-1]
			popIndex++
		}
	}
	return len(stack) == 0
}

// @lc code=end

