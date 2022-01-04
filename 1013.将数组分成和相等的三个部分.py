#
# @lc app=leetcode.cn id=1013 lang=python3
#
# [1013] 将数组分成和相等的三个部分
#

# @lc code=start
class Solution:
    def canThreePartsEqualSum(self, arr: List[int]) -> bool:
        # 双指针确定三个区块计算三个和比较即可
        # arr_length = len(arr)
        # for i in range(arr_length-1):
        #     for h in range(arr_length-1,-1,-1):
        #         if h <= i + 1:
        #             continue
        #         if sum(arr[:i+1]) == sum(arr[i+1:h]) == sum(arr[h:]):
        #             return True
        # return False
        sum_arr = sum(arr)
        if sum_arr % 3 != 0:
            return False
        single_target = sum_arr // 3
        length, i, curr = len(arr), 0, 0
        while i < length:
            curr += arr[i]
            if curr == single_target:
                break
            i += 1
        if curr != single_target:
            return False
        j = i + 1
        while j + 1 < length:
            curr += arr[j]
            if curr == 2*single_target:
                return True
            j += 1
        return False
# @lc code=end
