/*
 * @lc app=leetcode.cn id=1775 lang=golang
 *
 * [1775] 通过最少操作次数使数组的和相等
 */

// @lc code=start
func minOperations(nums1 []int, nums2 []int) (res int) {
	difference := 0
	allMap := [6]int{}
	// 得出差异值，并记录数字出现次数
	// 注意：我们只关心 大===>小  小===>大，不关心具体修改，所以实际上对1和5操作是一样的，所以1---5，2---4，3---3，4---2，5---1一一对应，从大到小去操作，到5时，不关心哪边，只关心1和5的修改使得一边变大一边变小，差值减小至0即为目标
	for _, num := range nums1 {
		difference += num
		allMap[num-1]++
	}
	for _, num := range nums2 {
		difference -= num
		allMap[5-(num-1)]++
	}
	// 转成正数，好操作
	if difference < 0 {
		difference = -difference
		for left, right := 0, 5; left < right; left, right = left+1, right-1 {
			allMap[left], allMap[right] = allMap[right], allMap[left] // 交换
		}
	}
	for i := 5; i >= 1; i-- { // 反向，从较大的数字开始，减少操作次数
		if allMap[i]*i >= difference { // 该数次可操作次数内可以达成目标
			// 当前次数加上当前差异值所需操作次数，比如差异值11和15（5三次）超出的部分需要向上取整即可，反正也不是真的取完
			return res + int(math.Ceil(float64(difference)/float64(i)))
		}
		// 不可达成目标则减去当前操作次数内修改最大值
		difference -= allMap[i] * i
		// 操作次数加上这个数字的操作次数
		res += allMap[i]
	}
	return -1
}

// @lc code=end

