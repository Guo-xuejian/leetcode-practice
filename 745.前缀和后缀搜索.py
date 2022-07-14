#
# @lc app=leetcode.cn id=745 lang=python3
#
# [745] 前缀和后缀搜索
#

# @lc code=start
class WordFilter:

    def __init__(self, words: List[str]):
        self.d = {}
        for i, word in enumerate(words):
            m = len(word)
            for prefixLength in range(1, m + 1):
                for suffixLength in range(1, m + 1):
                    self.d[word[:prefixLength] + '#' + word[-suffixLength:]] = i


    def f(self, pref: str, suff: str) -> int:
        return self.d.get(pref + '#' + suff, -1)



# Your WordFilter object will be instantiated and called as such:
# obj = WordFilter(words)
# param_1 = obj.f(pref,suff)
# @lc code=end

