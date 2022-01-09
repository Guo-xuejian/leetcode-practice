/*
 * @lc app=leetcode.cn id=1629 lang=golang
 *
 * [1629] 按键持续时间最长的键
 */

// @lc code=start
func slowestKey(releaseTimes []int, keysPressed string) byte {
	res := keysPressed[0]
	maxTime := releaseTimes[0]
	for i := 1; i < len(keysPressed); i++ {
		key := keysPressed[i]
		time := releaseTimes[i] - releaseTimes[i-1]
		if time > maxTime || time == maxTime && key > res {
			res = key
			maxTime = time
		}
	}
	return res
}

// @lc code=end

