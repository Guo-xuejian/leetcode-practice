# 剑指 Offer II 083. 没有重复元素集合的全排列
# 给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。

# 示例 1：
# 输入：nums = [1,2,3]
# 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
# 示例 2：
# 输入：nums = [0,1]
# 输出：[[0,1],[1,0]]

# 示例 3：
# 输入：nums = [1]
# 输出：[[1]]

# 提示：
# 1 <= nums.length <= 6
# -10 <= nums[i] <= 10
# nums 中的所有整数 互不相同

# 注意：本题与主站 46 题相同：https://leetcode-cn.com/problems/permutations/


class Solution:
    # 执行用时：32 ms, 在所有 Python3 提交中击败了78.20%的用户
    # 内存消耗：15.1 MB, 在所有 Python3 提交中击败了64.25%的用户
    def permute(self, nums: List[int]) -> List[List[int]]:
        length = len(nums)
        res = []

        def dfs(num_list):  # 深度优先搜索
            if len(num_list) == length:  # 长度满足要求时写入结果集
                res.append([num for num in num_list])   # 记住重新写入而不是直接传入
            else:
                for one in nums:
                    if one not in num_list:  # 长度不满足要求时写入不同数字
                        num_list.append(one)
                        dfs(num_list)
                        num_list.remove(one)    # 不影响下一次循环，删除
        for one in nums:
            dfs([one])  # 分别作为起点
        return res
