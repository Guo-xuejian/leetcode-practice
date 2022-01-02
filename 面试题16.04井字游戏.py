# 面试题 16.04. 井字游戏
# 设计一个算法，判断玩家是否赢了井字游戏。输入是一个 N x N 的数组棋盘，由字符" "，"X"和"O"组成，其中字符" "代表一个空位。

# 以下是井字游戏的规则：

# 玩家轮流将字符放入空位（" "）中。
# 第一个玩家总是放字符"O"，且第二个玩家总是放字符"X"。
# "X"和"O"只允许放置在空位中，不允许对已放有字符的位置进行填充。
# 当有N个相同（且非空）的字符填充任何行、列或对角线时，游戏结束，对应该字符的玩家获胜。
# 当所有位置非空时，也算为游戏结束。
# 如果游戏结束，玩家不允许再放置字符。
# 如果游戏存在获胜者，就返回该游戏的获胜者使用的字符（"X"或"O"）；如果游戏以平局结束，则返回 "Draw"；如果仍会有行动（游戏未结束），则返回 "Pending"


class Solution:
    def tictactoe(self, board: List[str]) -> str:
        def win(c):
            cn = c * n  # 根据传入参数计算出获胜情况
            return any([
                any(row == cn for row in board),    # 行
                any(''.join(col) == cn for col in zip(*board)),     # 列
                all(board[i][i] == c for i in range(n)),        # 对角
                all(board[i][n-i-1] == c for i in range(n))]    # 斜对角
            )
        n = len(board)
        if win('X'):    # X 一方获胜判定
            return 'X'
        if win('O'):    # O 一方获胜判定
            return 'O'
        if any(' ' in row for row in board):    # 仍有空位则继续
            return 'Pending'
        return 'Draw'   # 平局

