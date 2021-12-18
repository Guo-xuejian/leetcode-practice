/*
 * @lc app=leetcode.cn id=1333 lang=golang
 *
 * [1333] 餐厅过滤器
 */

// @lc code=start
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) (res []int) {
	for index := range restaurants {
		// 满足三个过滤器
		if veganFriendly == (veganFriendly&restaurants[index][2]) && restaurants[index][3] <= maxPrice && restaurants[index][4] <= maxDistance {
			res = append(res, index)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		// 按照 rating 排序，从高到低，相同则按照 id 从高到低排序
		if restaurants[res[i]][1] != restaurants[res[j]][1] {
			return restaurants[res[i]][1] > restaurants[res[j]][1]
		}
		return restaurants[res[i]][0] > restaurants[res[j]][0]
	})
	for idx := range res {
		index := res[idx]
		res[idx] = restaurants[index][0]
	}
	return
}

// @lc code=end

