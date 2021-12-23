# 实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

# 示例 1：
# 输入: s = "leetcode"
# 输出: false 

# 示例 2：
# 输入: s = "abc"
# 输出: true


class Solution:
    def isUnique(self, astr: str) -> bool:
        return len(astr) == len(set(astr))