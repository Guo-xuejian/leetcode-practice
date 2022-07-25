/*
 * @lc app=leetcode.cn id=1128 lang=golang
 *
 * [1128] 等价多米诺骨牌对的数量
 */

// @lc code=start
package main

func numEquivDominoPairs(dominoes [][]int) (res int) {
	num := [100]int{}
	val := 0
	dominoesLength := len(dominoes)
	for i := 0; i < dominoesLength; i++ {
		if dominoes[i][0] <= dominoes[i][1] {
			val = dominoes[i][0]*10 + dominoes[i][1]
		} else {
			val = dominoes[i][0] + dominoes[i][1]*10
		}
		res += num[val]
		num[val]++
	}
	return
}

// @lc code=end
