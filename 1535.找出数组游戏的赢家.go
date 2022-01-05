/*
 * @lc app=leetcode.cn id=1535 lang=golang
 *
 * [1535] 找出数组游戏的赢家
 */

// @lc code=start
func getWinner(arr []int, k int) int {
	prev := max(arr[0], arr[1])
	if k == 1 { // 只需获胜一次的话也只需要判定前两个元素
		return prev
	}

	constantWining := 1 // 持续获胜次数
	arrLength := len(arr)

	for i := 2; i < arrLength; i++ { // 第三个元素开始
		if prev > arr[i] { // 持续获胜，增加获胜次数，达到退出条件退出
			constantWining++
			if constantWining == k {
				return prev
			}
		} else { // 出现更大值，覆盖，获胜次数置为 1
			prev = arr[i]
			constantWining = 1
		}
	}
	// 一轮遍历之后出现最大值，也就无需后续的比较
	return prev
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

