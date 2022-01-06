/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	names := strings.Split(path, "/") // 切分之后按情况讨论
	stack := []string{}
	for _, name := range names {
		if name == ".." { // 上一级目录则去除
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if len(name) > 0 && name != "." { // 正常目录添加，当前目录不做处理
			stack = append(stack, name)
		}
	}
	// 返回值以根目录开始，结尾无 /
	return "/" + strings.Join(stack, "/")
}

// @lc code=end

