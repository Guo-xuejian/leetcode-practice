/*
 * @lc app=leetcode.cn id=819 lang=golang
 *
 * [819] 最常见的单词
 */

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {
    ban := map[string]bool{}
    for _, s := range banned {
        ban[s] = true
    }
    freq := map[string]int{}
    maxFreq := 0
    var word []byte
    for i, n := 0, len(paragraph); i <= n; i++ {
        if i < n && unicode.IsLetter(rune(paragraph[i])) {
            word = append(word, byte(unicode.ToLower(rune(paragraph[i]))))
        } else if word != nil {
            s := string(word)
            if !ban[s] {
                freq[s]++
                maxFreq = max(maxFreq, freq[s])
            }
            word = nil
        }
    }
    for s, f := range freq {
        if f == maxFreq {
            return s
        }
    }
    return ""
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

// @lc code=end

