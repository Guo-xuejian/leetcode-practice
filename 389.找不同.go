/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了77.60%的用户
func findTheDifference(s string, t string) (res byte) {
	if len(s) == 0 {
		res = t[0]
		return
	}
	letterMap := map[byte]int{}
	index := 0
	for index < len(s) {
		// 互相抵消
		letterMap[s[index]]++
		letterMap[t[index]]--
		index++
	}
	letterMap[t[len(t)-1]]-- // t 长度多 1
	for key, val := range letterMap {
		if val == -1 { // 计数值为 -1 的就是添加的字母
			res = key
			break
		}
	}
	return
}

// @lc code=end

