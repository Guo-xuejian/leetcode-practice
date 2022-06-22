/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */

// @lc code=start
func findSubstring(s string, words []string) (ans []int) {
    ls, m, n := len(s), len(words), len(words[0])
    for i := 0; i < n && i+m*n <= ls; i++ {
        differ := map[string]int{}
        for j := 0; j < m; j++ {
            differ[s[i+j*n:i+(j+1)*n]]++
        }
        for _, word := range words {
            differ[word]--
            if differ[word] == 0 {
                delete(differ, word)
            }
        }
        for start := i; start < ls-m*n+1; start += n {
            if start != i {
                word := s[start+(m-1)*n : start+m*n]
                differ[word]++
                if differ[word] == 0 {
                    delete(differ, word)
                }
                word = s[start-n : start]
                differ[word]--
                if differ[word] == 0 {
                    delete(differ, word)
                }
            }
            if len(differ) == 0 {
                ans = append(ans, start)
            }
        }
    }
    return
}

// @lc code=end

