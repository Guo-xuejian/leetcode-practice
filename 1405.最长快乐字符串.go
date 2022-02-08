/*
 * @lc app=leetcode.cn id=1405 lang=golang
 *
 * [1405] 最长快乐字符串
 */

// @lc code=start
func longestDiverseString(a, b, c int) string {
	ans := []byte{}
	cnt := []struct {
		c  int
		ch byte
	}{{a, 'a'}, {b, 'b'}, {c, 'c'}}
	for {
		sort.Slice(cnt, func(i, j int) bool { return cnt[i].c > cnt[j].c })
		hasNext := false
		for i, p := range cnt {
			if p.c <= 0 {
				break
			}
			m := len(ans)
			if m >= 2 && ans[m-2] == p.ch && ans[m-1] == p.ch {
				continue
			}
			hasNext = true
			ans = append(ans, p.ch)
			cnt[i].c--
			break
		}
		if !hasNext {
			return string(ans)
		}
	}
}

// @lc code=end

