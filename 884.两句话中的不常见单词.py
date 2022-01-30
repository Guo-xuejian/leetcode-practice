#
# @lc app=leetcode.cn id=884 lang=python3
#
# [884] 两句话中的不常见单词
#

# @lc code=start
class Solution:
    def uncommonFromSentences(self, s1: str, s2: str) -> List[str]:
        # 满足条件的单词一定是在两个句子加总起来只出现一次
        # 那么就都写入一个字典，通过字典键值来判断是否只出现一次
        word_times_dict = {}
        for word in s1.split() + s2.split():
            if word not in word_times_dict:
                word_times_dict[word] = 1
            else:
                word_times_dict[word] += 1
        # 写入只出现一次的单词
        return [word for word, times in word_times_dict.items() if times == 1]
        # @lc code=end

