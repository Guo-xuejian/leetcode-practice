#
# @lc app=leetcode.cn id=139 lang=python3
#
# [139] 单词拆分
#

# @lc code=start
class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        word_dict = {}
        for word in wordDict:   # 字典加速访问
            word_dict[word] = True
        dp = [False for _ in range(len(s)+1)]
        dp[0] = True

        for i in range(1, len(s) + 1):
            for j in range(i):
                # s[:j] 存在同时 s[j:i] 存在，表示当前部分可以由字典中字符串组成
                if dp[j] and s[j:i] in word_dict:
                    dp[i] = True
                    break
        return dp[len(s)]
# @lc code=end

