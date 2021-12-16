# 正整数 n 代表生成括号的对数，请设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
# 示例 1：
# 输入：n = 3
# 输出：["((()))","(()())","(())()","()(())","()()()"]

# 示例 2：
# 输入：n = 1
# 输出：["()"]

# 提示：
# 1 <= n <= 8

class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []
        if n <= 0:
            return res

        def dfs(path, open_num, close_num) -> None:
            # 括号数量需要符合要求
            if open_num > n or close_num > open_num:
                return

            if len(path) == 2 * n:
                res.append(path)
                return

            dfs(path+"(", open_num + 1, close_num)
            dfs(path+")", open_num, close_num + 1)
        dfs("", 0, 0)
        return res
