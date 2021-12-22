/*
 * @lc app=leetcode.cn id=1154 lang=golang
 *
 * [1154] 一年中的第几天
 */

// @lc code=start
func dayOfYear(date string) (res int) {
	splitString := strings.Split(date, "-")
	year, _ := strconv.Atoi(splitString[0])
	month, _ := strconv.Atoi(splitString[1])
	day, _ := strconv.Atoi(splitString[2])

	daysEachMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		daysEachMonth[1]++
	}
	for i := 0; i < month-1; i++ {
		res += daysEachMonth[i]
	}
	res += day
	return
}

// @lc code=end

