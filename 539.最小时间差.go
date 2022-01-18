/*
 * @lc app=leetcode.cn id=539 lang=golang
 *
 * [539] 最小时间差
 */

// @lc code=start
func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	minTime := math.MaxInt32
	timeLap0 := getMinutes(timePoints[0])
	timeLapPre := timeLap0
	timeNow := 0
	for i := 1; i < len(timePoints); i++ { // 从最新小时间开始逐一比较，首位和末尾有可能跨天，后面处理
		timeNow = getMinutes(timePoints[i])
		timeLap := timeNow - timeLapPre
		timeLapPre = timeNow
		if timeLap < minTime {
			minTime = timeLap
			if minTime == 0 {
				return minTime
			}
		}
	}
	// 处理跨天,，其实也就是末尾和首位比较差值即可，但是00：00 这样的时间需要加上 1440（24*60）分钟后再比较
	if timeLap0+1440-timeNow < minTime {
		return timeLap0 + 1440 - timeNow
	}
	return minTime
}

func getMinutes(t string) int {
	return (int((t[0]-'0'))*10+int(t[1]-'0'))*60 + (int((t[3]-'0'))*10 + int(t[4]-'0'))
}

// @lc code=end

