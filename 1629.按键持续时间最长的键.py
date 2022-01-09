#
# @lc app=leetcode.cn id=1629 lang=python3
#
# [1629] 按键持续时间最长的键
#

# @lc code=start
class Solution:
    def slowestKey(self, releaseTimes: List[int], keysPressed: str) -> str:
        res = keysPressed[0]    # 当前持续时间最长按键
        max_time = releaseTimes[0]  # 当前持续最长时间
        for i in range(1, len(releaseTimes)):
            time_last = releaseTimes[i] - releaseTimes[i-1]
            # 大于
            # 等于但是需要字母序大于 res
            # if time_last > max_time or (time_last == max_time and ord(keysPressed[i]) > ord(res)):
            if time_last > max_time or (time_last == max_time and keysPressed[i] > res):
                max_time = time_last
                res = keysPressed[i]
        return res
# @lc code=end

