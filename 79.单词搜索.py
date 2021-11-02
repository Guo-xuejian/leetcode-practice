#
# @lc app=leetcode.cn id=79 lang=python3
#
# [79] 单词搜索
#

# @lc code=start
class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        len1 = len(board)
        len2 = len(board[0])
        len3 = len(word)
        # 深度优先所有====所有情况都需要判定
        def dfs(i,j,k):
            # 数组不能越界
            if not 0 <= i < len1 or not 0 <= j < len2 or board[i][j] != word[k]:
                return False
            # 当前字符是符合要求的，判定是否为最后一个元素，如果是则返回
            if k == len3 - 1:
                return True
            board[i][j] = ""
            res = dfs(i-1,j,k+1) or dfs(i+1,j,k+1) or dfs(i,j-1,k+1) or dfs(i,j+1,k+1)
            board[i][j] = word[k]
            return res
        # 遍历所有的元素并将其最为起点开始判定
        for i in range(len1):
            for j in range(len2):
                if dfs(i,j,0):
                    return True
        return False
# @lc code=end

