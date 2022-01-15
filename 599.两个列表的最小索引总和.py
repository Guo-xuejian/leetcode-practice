#
# @lc app=leetcode.cn id=599 lang=python3
#
# [599] 两个列表的最小索引总和
#

# @lc code=start
class Solution:
    def findRestaurant(self, list1: List[str], list2: List[str]) -> List[str]:
        list1_length, list2_length = len(list1), len(list2)
        min_index_sum = list1_length + list2_length
        # 字典
        indexSum_stringArray_dict = {}
        for i in range(list1_length):
            for j in range(list2_length):
                curr_index_sum = i + j
                if curr_index_sum > min_index_sum:
                    # 大于时，后续也不会满足，且大于的结果不会是最终结果
                    break
                if list1[i] == list2[j]:
                    if curr_index_sum in indexSum_stringArray_dict:
                        # 已存在 添加
                        indexSum_stringArray_dict[curr_index_sum].append(list1[i])
                    else:
                        # 不存在 写入
                        indexSum_stringArray_dict[curr_index_sum] = [list1[i]]
                    # 比较得出最小索引和
                    if curr_index_sum < min_index_sum:
                        min_index_sum = curr_index_sum
        # 根据最小索引返回商店列表
        return indexSum_stringArray_dict[min_index_sum]
# @lc code=end

