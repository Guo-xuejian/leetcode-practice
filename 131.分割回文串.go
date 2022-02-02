/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) (res [][]string) {
	path := []string{}
	backtrack(s, &res, &path)
	return
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func backtrack(s string, res *[][]string, path *[]string) {
	if len(s) < 1 {
		*res = append(*res, *path)
		return
	}

	for i := 1; i < len(s)+1; i++ {
		if isPalindrome(s[:i]) {
			temp := make([]string, len(*path))
			copy(temp, *path)
			temp = append(temp, s[:i])
			backtrack(s[i:], res, &temp)
		}
	}
}

// @lc code=end

