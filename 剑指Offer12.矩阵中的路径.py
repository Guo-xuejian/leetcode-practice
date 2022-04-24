# 剑指 Offer 12. 矩阵中的路径
# 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

# 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

 

# 例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。



 

# 示例 1：

# 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
# 输出：true
# 示例 2：

# 输入：board = [["a","b"],["c","d"]], word = "abcd"
# 输出：false
 

# 提示：

# 1 <= board.length <= 200
# 1 <= board[i].length <= 200
# board 和 word 仅由大小写英文字母组成


class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        def dfs(i, j, k):
            # 数组不能越界，同时字符要符合要求
            if not 0 <= i < len(board) or not 0 <= j < len(board[0]) or board[i][j] != word[k]: return False
            if k == len(word) - 1: return True      # 最后一个元素则直接返回
            board[i][j] = ''    # 标记为已访问，这一步是为了下面的各个分支走不同的路线而写
            # 所有的子节点判定，只要有一条路走通，那就是 True
            res = dfs(i + 1, j, k + 1) or dfs(i - 1, j, k + 1) or dfs(i, j + 1, k + 1) or dfs(i, j - 1, k + 1)
            board[i][j] = word[k]   # 路线走完，将对应的值返回
            return res

        for i in range(len(board)):
            for j in range(len(board[0])):
                # 遍历每一个元素并将其作为起点判定是否符合要求
                if dfs(i, j, 0): return True
        return False    # 都不符合要求则 False
