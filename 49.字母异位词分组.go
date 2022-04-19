/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	allRes := map[string][]string{}
	for _, word := range strs {
		currSlice := []byte(word)
		sort.Slice(currSlice, func(i, j int) bool {
			return currSlice[i] < currSlice[j]
		})
		key := string(currSlice)
		if _, ok := allRes[key]; ok {
			allRes[key] = append(allRes[key], word)
		} else {
			allRes[key] = []string{word}
		}
	}
	res := make([][]string, 0, len(allRes))
	for _, val := range allRes {
		res = append(res, val)
	}
	return res
}
// @lc code=end

