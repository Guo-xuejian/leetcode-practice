// 剑指 Offer II 004. 只出现一次的数字
// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

// 示例 1：
// 输入：nums = [2,2,3,2]
// 输出：3

// 示例 2：
// 输入：nums = [0,1,0,1,0,1,100]
// 输出：100

// 提示：
// 1 <= nums.length <= 3 * 104
// -231 <= nums[i] <= 231 - 1
// nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次

func singleNumber(nums []int) (res int) {
	// 先用 map 实现
	/*
	   执行用时：4 ms, 在所有 Go 提交中击败了96.62%的用户
	   内存消耗：4.1 MB, 在所有 Go 提交中击败了21.97%的用户
	*/
	// numTimesMap := map[int]int{}
	// for _, num := range nums {
	//     numTimesMap[num]++
	// }
	// for key, val := range numTimesMap {
	//     if val == 1 {
	//         res = key
	//         break
	//     }
	// }
	// return

	// 排序实现
	/*
	   执行用时：8 ms, 在所有 Go 提交中击败了12.11%的用户
	   内存消耗：3.4 MB, 在所有 Go 提交中击败了100.00%的用户
	*/
	sort.Ints(nums)
	preNum := 1
	preIndex := 0
	// 从索引 1 开始
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] { // 不同
			if preNum == 1 { // 之前的数字只出现一次满足条件
				return nums[preIndex]
			} else { // 不满足重置出现次数和索引
				preNum = 1
				preIndex = i
			}
		} else { // 相同次数增加
			preNum++
		}
	}
	// 到这里意味着最后一个数字才是答案，直接返回即可
	return nums[preIndex]
}