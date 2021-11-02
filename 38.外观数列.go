/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	prev := "1"
	for i := 0; i < n-1; i++ {
		curr := &strings.Builder{}
		pos := 0
		start := 0
		for pos < len(prev) {
			for pos < len(prev) && prev[pos] == prev[start] {
				// 相同数字 pos 计数
				pos++
			}
			// 不同数字中断并给出一部分结果
			curr.WriteString(strconv.Itoa(pos - start))
			curr.WriteByte(prev[start])
			// 设置新的起点，继续
			start = pos
		}
		// 一轮遍历结束覆盖 prev，作为下一轮遍历的字符串
		prev = curr.String()
	}
	return prev
}

// @lc code=end

