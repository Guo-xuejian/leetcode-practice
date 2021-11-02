#
# @lc app=leetcode.cn id=5 lang=python3
#
# [5] 最长回文子串
#

# @lc code=start
class Solution:
    def longestPalindrome(self, s: str) -> str:
        # 从中心出发，从每一个点开始，当作中心，向外扩散
        if not s:
            return s
        # start ==> 最终的左边界  end ==> 有边界（但是注意数组下标需要最终 +1）
        # left right 分别对应单中心和双中心遍历类型
        start = end = left1 = right1 = left2 = right2 = 0
        for i in range(len(s)):
            left1, right1 = self.expandAroundCenter(s, i, i)    # 但中心
            left2, right2 = self.expandAroundCenter(s, i, i+1)  #双中心
            if right1 - left1 > end - start:
                start, end = left1, right1
            if right2 - left2 > end - start:
                start, end = left2, right2
        return s[start : end + 1]

    def expandAroundCenter(self, s, x, y):
        while x >= 0 and y < len(s) and s[x] == s[y]:
            x -= 1
            y += 1
        return x + 1, y - 1
# @lc code=end

