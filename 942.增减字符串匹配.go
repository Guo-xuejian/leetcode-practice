/*
 * @lc app=leetcode.cn id=942 lang=golang
 *
 * [942] 增减字符串匹配
 */

// @lc code=start
func diStringMatch(s string) (res []int) {
	minNum, maxNum := 0, len(s)
	for _, letter := range s {
		if letter == 'I' {
			res = append(res, minNum)
			minNum++
		} else {
			res = append(res, maxNum)
			maxNum--
		}
	}
	res = append(res, minNum)
	return
}

// 2022-05-09
// func diStringMatch(s string) []int {
//     left, right, ans := 0, len(s), make([]int, len(s)+1)
//     for i:= range s{
//         if s[i] == 'I'{
//             ans[i] = left
//             left++
//         }else{
//             ans[i] = right
//             right--
//         }
//     }
//     ans[len(s)] = left

//     return ans
// }
// @lc code=end

