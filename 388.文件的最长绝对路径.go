/*
 * @lc app=leetcode.cn id=388 lang=golang
 *
 * [388] 文件的最长绝对路径
 */

// @lc code=start
func lengthLongestPath(input string) (res int) {
	nodes := strings.Split(input, "\n")
	length := []int{0}
	for i := range nodes {
		t, l := getTab(nodes[i])
		if t > 0 {
			l += length[t-1]
		}
		if isFile(nodes[i]) {
			res = max(res, l+t)
		}
		if t >= len(length) {
			length = append(length, l)
		} else {
			length[t] = l
		}
	}
	return
}
func isFile(str string) bool {
	return strings.Contains(str, ".")
}
func getTab(str string) (t, l int) {
	for i := range str {
		if str[i] == '\t' {
			t++
		} else {
			break
		}
	}
	return t, len(str) - t
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

