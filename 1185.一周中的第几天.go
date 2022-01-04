/*
 * @lc app=leetcode.cn id=1185 lang=golang
 *
 * [1185] 一周中的第几天
 */

// @lc code=start
func dayOfTheWeek(day int, month int, year int) string {
	var month_day = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var day_week = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	var cur int = 0
	for y := 1971; y < year; y++ {
		cur += 365
		if is_leap_year(y) == true {
			cur++
		}
	}
	for m := 1; m < month; m++ {
		cur += month_day[m]
		if m == 2 && is_leap_year(year) == true {
			cur++
		}
	}
	cur += day
	// 1971.01.01起始点是周五，day_week从周天开始，所以加一个位移4
	return day_week[(4+cur)%7]
}

func is_leap_year(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// @lc code=end

