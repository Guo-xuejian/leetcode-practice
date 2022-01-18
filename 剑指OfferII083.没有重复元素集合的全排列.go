// 剑指 Offer II 083. 没有重复元素集合的全排列
// 给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。

// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 示例 2：
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]

// 示例 3：
// 输入：nums = [1]
// 输出：[[1]]

// 提示：
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同

// 注意：本题与主站 46 题相同：https://leetcode-cn.com/problems/permutations/

func permute(nums []int) (res [][]int) {
	length := len(nums)
	for _, v := range nums {
		dfs(&[]int{v}, &res, length, &nums) // 分别作为起点
	}
	return
}

func dfs(numSlice *[]int, res *[][]int, length int, nums *[]int) {
	if len(*numSlice) == length { // 长度符合要求写入结果
		// 注意引用传递的问题，创建个新切片写入
		newSlice := make([]int, len(*numSlice))
		copy(newSlice, *numSlice)
		*res = append(*res, newSlice)
	} else {
		for _, v := range *nums {
			status := false
			for _, val := range *numSlice { // 当前切片中不存在该数字
				if val == v {
					status = true
					break
				}
			}
			if !status { // 确定不存在
				*numSlice = append(*numSlice, v)           // 写入
				dfs(numSlice, res, length, nums)           // 递归
				*numSlice = (*numSlice)[:len(*numSlice)-1] // 不影响下一次循环，删除该元素
			}
		}
	}

}