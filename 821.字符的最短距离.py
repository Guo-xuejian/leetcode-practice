#
# @lc app=leetcode.cn id=821 lang=python3
#
# [821] 字符的最短距离
#

# @lc code=start
class Solution:
    def shortestToChar(self, s: str, c: str) -> List[int]:
        m = len(s)
        def getDistance(index: int) -> int:
            if s[index] == c:
                return 0
            left, right = index, index
            while left >= 0 or right < m:
                if left >= 0 and s[left] == c:
                    return abs(index - left)
                elif left >= 0:
                    left -= 1
                if right < m and s[right] == c:
                    return abs(index - right)
                elif right < m:
                    right += 1
            return -1
        res = [getDistance(i) for i in range(m)]
        return res
# @lc code=end

