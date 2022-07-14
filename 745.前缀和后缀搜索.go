/*
 * @lc app=leetcode.cn id=745 lang=golang
 *
 * [745] 前缀和后缀搜索
 */

// @lc code=start
type WordFilter map[string]int

func Constructor(words []string) WordFilter {
	wf := WordFilter{}
	for i, word := range words {
		for j, n := 1, len(word); j <= n; j++ {
			for k := 0; k < n; k++ {
				wf[word[:j]+"#"+word[k:]] = i
			}
		}
	}
	return wf
}

func (wf WordFilter) F(pref, suff string) int {
	if i, ok := wf[pref+"#"+suff]; ok {
		return i
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
// @lc code=end

