#
# @lc app=leetcode.cn id=1535 lang=python3
#
# [1535] 找出数组游戏的赢家
#

# @lc code=start
class Solution:
    def getWinner(self, arr: List[int], k: int) -> int:
        prev = max(arr[0], arr[1])
        if k == 1:  # 只需获胜一次则直接返回即可
            return prev

        constant_wining = 1     # 连续获胜次数
        arr_length = len(arr)

        # 只需要一轮遍历，获取到最大值之后即使没有满足 k 要求，后续比较也一定会满足
        for i in range(2, arr_length):
            curr = arr[i]
            if prev > curr:     # 持续获胜，获胜次数 +1，满足退出条件时退出
                constant_wining += 1
                if constant_wining == k:
                    return prev
            else:               # 出现更大值，获胜次数置为 1，重置较大值 prev
                constant_wining = 1
                prev = curr
        # 一轮之后就不用继续比较
        return prev
# @lc code=end
