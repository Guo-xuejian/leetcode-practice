/*
 * @lc app=leetcode.cn id=838 lang=golang
 *
 * [838] 推多米诺
 */

// @lc code=start
func pushDominoes(dominoes string) string {
	length := len(dominoes)
	res := make([]byte, length)
	for idx, letter := range dominoes {
		res[idx] = byte(letter)
	}
	i, left := 0, byte('L')

	for i < length {
		j := i
		var right byte
		for j < length && res[j] == '.' {
			j++
		}
		if j < length {
			right = res[j]
		} else {
			right = 'R'
		}
		if left == right {
			for i < j {
				res[i] = right
				i++
			}
		} else if left == 'R' && right == 'L' {
			k := j - 1
			for i < k {
				res[i] = 'R'
				res[k] = 'L'
				i++
				k--
			}
		}
		left = right
		i = j + 1
	}
	return string(res)
}

// @lc code=end

