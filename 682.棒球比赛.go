/*
 * @lc app=leetcode.cn id=682 lang=golang
 *
 * [682] 棒球比赛
 */

// @lc code=start
func calPoints(ops []string) (res int) {
	arr := []int{}
	for _, one := range ops {
		if one == "+" {
			arr = append(arr, arr[len(arr)-1]+arr[len(arr)-2])
		} else if one == "D" {
			arr = append(arr, arr[len(arr)-1]*2)
		} else if one == "C" {
			res -= arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			continue
		} else {
			num, _ := strconv.Atoi(one)
			arr = append(arr, num)
		}
		res += arr[len(arr)-1]
	}
	return
}

// @lc code=end

