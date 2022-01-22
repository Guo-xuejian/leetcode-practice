// 剑指 Offer II 092. 翻转字符
// 如果一个由 '0' 和 '1' 组成的字符串，是以一些 '0'（可能没有 '0'）后面跟着一些 '1'（也可能没有 '1'）的形式组成的，那么该字符串是 单调递增 的。

// 我们给出一个由字符 '0' 和 '1' 组成的字符串 s，我们可以将任何 '0' 翻转为 '1' 或者将 '1' 翻转为 '0'。

// 返回使 s 单调递增 的最小翻转次数。

// 示例 1：
// 输入：s = "00110"
// 输出：1
// 解释：我们翻转最后一位得到 00111.

// 示例 2：
// 输入：s = "010110"
// 输出：2
// 解释：我们翻转得到 011111，或者是 000111。

// 示例 3：
// 输入：s = "00011000"
// 输出：2
// 解释：我们翻转得到 00000000。

// 提示：
// 1 <= s.length <= 20000
// s 中只包含字符 '0' 和 '1'

func minFlipsMonoIncr(s string) int {
	length := len(s)
	preNum := make([]int, length+1)
	for idx, val := range s {
		if val == '1' { // 每一个起点之前的数字 1 的数量
			preNum[idx+1] = preNum[idx] + 1
		} else {
			preNum[idx+1] = preNum[idx]
		}
	}

	res := length
	for mid := 0; mid < length+1; mid++ { // 左侧有mid个数字，mid属于右侧部分
		L1 := preNum[mid]                                     // 每一个起点之前的数字 1 的数量，需要修改为 0
		R0 := (length - mid) - (preNum[length] - preNum[mid]) // 对应起点后面需要修改的 0 的数量
		res = min(res, L1+R0)                                 // 取得最小操作次数
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}