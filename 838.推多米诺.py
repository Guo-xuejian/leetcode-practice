#
# @lc app=leetcode.cn id=838 lang=python3
#
# [838] 推多米诺
#

# @lc code=start
class Solution:
    def pushDominoes(self, dominoes: str) -> str:
        s = list(dominoes)
        n, i, left = len(s), 0, 'L'
        while i < n:
            j = i
            while j < n and s[j] == '.':  # 找到一段连续的没有被推动的骨牌
                j += 1
            right = s[j] if j < n else 'R'
            if left == right:  # 方向相同，那么这些竖立骨牌也会倒向同一方向
                while i < j:
                    s[i] = right
                    i += 1
            elif left == 'R' and right == 'L':  # 方向相对，那么就从两侧向中间倒
                k = j - 1
                while i < k:
                    s[i] = 'R'
                    s[k] = 'L'
                    i += 1
                    k -= 1
            # left == 'L' right == 'R' 不需要推倒
            left = right    # 右边界成为新的左边界
            i = j + 1       # 回到一开始的样子
        return ''.join(s)

# @lc code=end
