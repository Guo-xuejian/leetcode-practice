/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) (res []string) {
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	// 内存消耗：1.9 MB, 在所有 Go 提交中击败了93.53%的用户
	wordMap := map[string]bool{} // 哈希表加快访问速度
	for _, word := range wordDict {
		wordMap[word] = true
	}
	check(s, &[]string{}, &res, &wordMap)
	return

}

func check(s string, tempSlice *[]string, res *[]string, wordMap *map[string]bool) {
	if len(s) < 1 { // 字符串遍历结束，说明完全匹配，拼接后写入结果
		*res = append(*res, strings.Join(*tempSlice, " "))
		return
	}

	for i := 0; i < len(s)+1; i++ { // 逐一作为切分点
		if (*wordMap)[s[:i]] { // 当前切分结果在字典中则继续
			// 做拷贝
			temp := make([]string, len(*tempSlice))
			copy(temp, *tempSlice)
			temp = append(temp, s[:i]) // 记录当前结果进入递归
			check(s[i:], &temp, res, wordMap)
		}
	}
}

// @lc code=end

