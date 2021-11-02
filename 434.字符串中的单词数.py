#
# @lc app=leetcode.cn id=434 lang=python3
#
# [434] 字符串中的单词数
#

# @lc code=start
class Solution:
    def countSegments(self, s: str) -> int:
        res = 0
        # 为避免重复计数，初始化 writeFlag 来表示是否仍然处在一个单词范围内
        # writeFlag 为真时放弃计数，表示还在该单词范围内
        writeFlag = False
        for one in s:
            if one == " ":
                writeFlag = False
                continue
            if not writeFlag:
                res += 1
                writeFlag = True
        return res
# @lc code=end

