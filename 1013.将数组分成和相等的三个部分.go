/*
 * @lc app=leetcode.cn id=1013 lang=golang
 *
 * [1013] 将数组分成和相等的三个部分
 */

// @lc code=start
package main

func canThreePartsEqualSum(arr []int) bool {
	var arrSum int
	for _, v := range arr { // 计算总和
		arrSum += v
	}
	if arrSum%3 != 0 { // 不能三等分则无结果
		return false
	}
	singleTarget := arrSum / 3 // 单份
	length, i, currSum := len(arr), 0, 0
	for i < length {
		currSum += arr[i]
		if currSum == singleTarget { // 第一区间找到结束索引
			break
		}
		i++
	}
	if currSum != singleTarget { // 可能结束了还是小于（也就是少一个结尾元素），这种情况不可能满足三区间
		return false
	}
	j := i + 1
	for j+1 < length { // 至少预留一个第三区间,所以再加一次 1
		currSum += arr[j]
		if currSum == 2*singleTarget { // 满足该条件即可，剩余部分（第三区间）之和必然为 singleTarget
			return true
		}
		j++
	}
	// 不存在第二个区间
	return false
}

// @lc code=end

