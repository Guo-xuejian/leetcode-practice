#
# @lc app=leetcode.cn id=969 lang=python3
#
# [969] 煎饼排序
#

# @lc code=start
class Solution:
    def pancakeSort(self, arr: List[int]) -> List[int]:
        res = []
        for n in range(len(arr), 1, -1):
            index = 0
            for i in range(n):
                if arr[i] > arr[index]:
                    index  = i
            
            # 如果已经是最后一位，没必要在移动
            if index == n - 1:
                continue
            m = index
            for i in range((m + 1) // 2):
                arr[i], arr[m - i] = arr[m - i], arr[i]
            
            for i in range(n // 2):
                arr[i], arr[n - 1- i] = arr[n - 1 - i], arr[i]
            res.append(index + 1)
            res.append(n)
        return res
# @lc code=end

