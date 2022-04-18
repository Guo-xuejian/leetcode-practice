# 给定一个不包含重复数字的数组，返回其全部子集，顺序无要求
#
# 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
#
# 输出指定输入的数组，列出该数组的所有子集合
# @param a int整型一维数组 假设a数组中的数字不相同
# @return int整型二维数组
#
class Solution:
    def SubSets(self , a: List[int]) -> List[List[int]]:
        # write code here
        # 动态规划，前缀和思想
        dp = [[a[-1]], []]
        for i in range(len(a) - 2, -1, -1):
            temp_list = [part + [a[i],] for part in dp]
            dp.extend(temp_list)
        return dp


# 一个 m*n 的网格，一个机器人从左上角出发，只能向右或者向下，请问一共有多少种走法从最左上角走到最右下角
#
# 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
#
# 输出网格的宽度和高度，返回共有多少种走法
# @param m int整型 网格的宽度
# @param n int整型 网格的高度
# @return int整型
#
class Solution:
    def PathsCnt(self , m: int, n: int) -> int:
        # write code here
        # 矩阵选手申请出战
        directions = [(1, 0), (0, 1)]
        grid = [[0 for _ in range(n)] for _ in range(m)]
        def dfs(i, j):
            # 越界判定，或者已经走过了
            if i >= m or j >= n:
                return
            grid[i][j] += 1
            for direction in directions:
                dfs(i + direction[0], j + direction[1])
        dfs(0, 0)
        return grid[m - 1][n - 1]
