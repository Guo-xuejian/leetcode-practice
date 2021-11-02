/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
import (
	"strconv"
)

func isPalindrome(x int) bool {
	// 转换为字符串
	stringX := strconv.Itoa(x)
	pre := 0
	back := 0
	lengthOfInt := len(stringX)
	// 如果是偶数，那么前后指针差距为1
	// 如果是奇数，那么前后指针初始值相等
	if lengthOfInt%2 == 0 {
		pre = lengthOfInt/2 - 1
		back = pre + 1
	} else {
		pre = lengthOfInt / 2
		back = lengthOfInt / 2
	}
	for {
		if stringX[pre] != stringX[back] {
			return false
		}
		pre--
		back++
		if pre < 0 {
			break
		}
	}
	return true

}

// @lc code=end

