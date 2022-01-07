/*
 * @lc app=leetcode.cn id=1652 lang=golang
 *
 * [1652] 拆炸弹
 */

// @lc code=start
func decrypt(code []int, k int) []int {
	codeLength := len(code)
	if k == 0 {
		for i := 0; i < codeLength; i++ {
			code[i] = 0
		}
		return code
	}
	res := []int{}
	if k > 0 {
		if k == 1 { // 实际上就是整体后移一位，注意零号元素应该放到后面
			code = append(code, code[0])
			return code[1:]
		} else if k == codeLength-1 { // 这种情况下对应位结果恰好等于 code 求和后减去该位数字
			allCodeSum := 0
			for _, num := range code {
				allCodeSum += num
			}
			for i := 0; i < codeLength; i++ {
				code[i] = allCodeSum - code[i]
			}
			return code
		} else {
			for i := 0; i < codeLength; i++ {
				currNum := 0
				index := 0
				if i+1 < codeLength { // 初始索引越界判定
					index = i + 1
				} else {
					index = 0
				}
				finalIndex := 0
				if i+k < codeLength { // 最终索引越界判定
					finalIndex = i + k
				} else {
					finalIndex = (i + k) % codeLength // 越界取余
				}
				for index != finalIndex {
					currNum += code[index]
					if index+1 < codeLength {
						index++
					} else {
						index = 0 // 内部越界回到零号索引
					}
				}
				currNum += code[index] // 由于没法通过 < 或者 > 作为循环，最后需要加上最终索引处数字
				res = append(res, currNum)
			}
		}
	} else {
		for i := 0; i < codeLength; i++ {
			currNum := 0
			index := 0
			if i-1 >= 0 {
				index = i - 1
			} else {
				index = codeLength - 1
			}

			finalIndex := 0
			if i+k >= 0 {
				finalIndex = i + k
			} else {
				finalIndex = codeLength + i + k // 越界从后开始
			}
			for index != finalIndex {
				currNum += code[index]
				if index-1 >= 0 {
					index--
				} else {
					index = codeLength - 1 // 内部越界去到末尾
				}
			}
			currNum += code[index]
			res = append(res, currNum)
		}
	}

	return res
}

// @lc code=end

