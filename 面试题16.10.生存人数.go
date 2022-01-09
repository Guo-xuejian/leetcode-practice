// 面试题 16.10. 生存人数
// 给定 N 个人的出生年份和死亡年份，第 i 个人的出生年份为 birth[i]，死亡年份为 death[i]，实现一个方法以计算生存人数最多的年份。
// 你可以假设所有人都出生于 1900 年至 2000 年（含 1900 和 2000 ）之间。如果一个人在某一年的任意时期处于生存状态，那么他应该被纳入那一年的统计中。例如，生于 1908 年、死于 1909 年的人应当被列入 1908 年和 1909 年的计数。
// 如果有多个年份生存人数相同且均为最大值，输出其中最小的年份。

// 示例：
// 输入：
// birth = {1900, 1901, 1950}
// death = {1948, 1951, 2000}
// 输出： 1901

// 提示：
// 0 < birth.length == death.length <= 10000
// birth[i] <= death[i]

func maxAliveYear(birth []int, death []int) int {
	// Go 的 map 和 Python 的 collections.Counter 一样好用，不存在的直接返回零值
	birthMap := make(map[int]int)
	deathMap := make(map[int]int)
	for _, one := range birth {
		birthMap[one]++
	}
	for _, one := range death {
		deathMap[one]++
	}
	liveNum, liveNumMax, liveMaxYear := 0, 0, 0
	for year := 1900; year <= 2000; year++ { // 必须逐年遍历，不然一些年份去世的人是无法被有效记录的
		liveNum += birthMap[year]
		if liveNum > liveNumMax { // 先计算人数，不需要等于，题目要求的就是最小的年份
			liveNumMax = liveNum
			liveMaxYear = year
		}
		liveNum -= deathMap[year] // 最后减去今年去世的，于是今年去世的就被计算在内，但是下一年不会
	}
	return liveMaxYear
}