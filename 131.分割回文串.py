#
# @lc app=leetcode.cn id=131 lang=python3
#
# [131] 分割回文串
#

# @lc code=start
class Solution(object):
    def partition(self, s):
        self.isPalindrome = lambda s : s == s[::-1] # 一个优美的函数
        res = []
        self.backtrack(s, res, [])
        return res
        
    def backtrack(self, s, res, path):
        if not s:   # 递归至空字符串结束
            res.append(path)
            return
        for i in range(1, len(s) + 1): #注意起始和结束位置
            if self.isPalindrome(s[:i]):    # 回文
                # path + [s[:i]] 加入当前回文子串再递归
                self.backtrack(s[i:], res, path + [s[:i]])
# @lc code=end

