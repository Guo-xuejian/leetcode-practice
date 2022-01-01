#
# @lc app=leetcode.cn id=2023 lang=python3
#
# [2023] 连接后等于目标字符串的字符串对
#

# @lc code=start
class Solution:
    def numOfPairs(self, nums: List[str], target: str) -> int:
        res = 0
        nums_length = len(nums)
        target_length = len(target)
        index_length_list = [0 for _ in range(nums_length)]
        for k, v in enumerate(nums):
            index_length_list[k] = [k, len(v)]
        # 根据长度来排个序
        index_length_list.sort(key=lambda x: (x[1]))
        for i in range(nums_length):
            curr_num = nums[index_length_list[i][0]]    # 当前字符
            curr_num_length = index_length_list[i][1]   # 当前字符长度
            if curr_num_length > target_length:
                break
            temp_length = target_length - curr_num_length   # 长度差值
            for j in range(0, i+1):
                if j == i:  # 不能和自己组合
                    continue
                # 当字符差值等于长度差值同时字符串正反组合等于最终要求则 res++
                if index_length_list[j][1] == temp_length:
                    if nums[index_length_list[j][0]] + curr_num == target:
                        res += 1
                    # 这种组合方式使得交换方向组合结果只有其一，需要额外处理，只要满足正反交换满足结果就再 res++ 一次
                    if curr_num + nums[index_length_list[j][0]] == target:
                        res += 1
        return res
# @lc code=end
