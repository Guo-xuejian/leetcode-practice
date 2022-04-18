# 红薯地
# 时间限制： 3000MS
# 内存限制： 589824KB
# 题目描述：
# 薯队长从北向南穿过一片红薯地（南北长 M 米，东西宽 N 米），红薯地被划分为 1米x1米 的方格，他可以从北边的任何一个格子出发，到达南边的任何一个格子，但每一步只能走到东南、正南、西南方向的三个格子之一，而且不能跨出红薯地，他可以获得经过的格子上的所有红薯，请问他可以获得最多的红薯个数。



# 输入描述
# 输入是 MxN + 2 个整数

# 第一个数字是 M， 1<=M<=1000

# 第二个数字是 N， 1<=N<=1000

# 后面是 MxN 个数字，代表 MxN 个方格上每个方格中红薯的个数 Amn，顺序是 A11 A12 ... A21 A22...

# 输出描述
# 输出是一个整数 X


# 样例输入
# 3 2
# 1 2 
# 3 1 
# 1 4
# 样例输出
# 9
#!/bin/python
# -*- coding: utf8 -*-
import sys
import os
import re

class Solution:
    def pathSum(self, array):
        # Write Code Here 
        m, n = len(array), len(array[0])
        self.res = 0
        directions = [(1, -1), (1, 0), (1, 1)]
        # 两边之和大于第三边，一直走东南，遇到东边界后走西南
        def dfs(i, j, curr):
            if i == m - 1:
                if curr + array[i][j] > self.res:
                    self.res = curr + array[i][j]
                    return
            for direction in directions:
                x, y = i + direction[0], j + direction[1]
                if x >= m or y < 0 or y >= n:
                    continue
                dfs(x, y, curr + array[i][j])

        for i in range(n):
            dfs(0, i, 0)
        return self.res
        

array_rows, array_cols = list(map(int, input().split()))

array = []
for array_i in range(array_rows):
    array.append(list(map(int, input().split())))

s = Solution()
res = s.pathSum(array)

print(res)





# 选礼物
# 时间限制： 3000MS
# 内存限制： 589824KB
# 题目描述：
# 薯队长最近在参加了一个活动，在活动结束后，可以挑选一些礼物留作纪念，主办方提供了N个礼物以供挑选，每个礼物有一个虚拟的价值，范围在0～10^9之间，薯队长可以从中挑选K个礼物（2 <= K <= N）。薯队长不想选价值过近的礼物，所以薯队长找到您帮忙，看从中选K个礼物，其中价值最接近的两件礼物之间相差值尽可能大。



# 输入描述
# 第一行两个整数N K

# 之后N行，每行一个整数，即每个礼物的价值

# 输出描述
# 一行，即选K个礼物中价值最接近的两件礼物之间相差值的最大值

# 比如样例中，从5个物品里选3个，选择价值为2、9、5的这3个礼物，价值最接近的两个礼物是2和5，相差3，是最大的一个结果


# 样例输入
# 5 3
# 2
# 3
# 9
# 5
# 10
# 样例输出
# 3

num_list = input().split(" ")
n, k = int(num_list[0]), int(num_list[1])
nums = [0 for _ in range(n)]
for i in range(n):
    curr = int(input())
    nums[i] = curr
nums.sort()
# 不要求最终价值最大，只要求 k 个礼物中价值最接近的两个差值最大











# 方块消除
# 时间限制： 3000MS
# 内存限制： 589824KB
# 题目描述：
# 薯队长最近在玩一个游戏，这个游戏桌上会有一排不同颜色的方块，每次薯队长可以选择一个方块，便可以消除这个方块以及和他左右相临的若干的颜色相同的方块，而每次消除的方块越多，得分越高。

# 具体来说，桌上有N个方块排成一排（1 <= N <= 200），每个方块有一个颜色，用1～N之间的一个整数表示，相同的数字代表相同的颜色，每次消除的时候，会把连续的K个相同颜色的方块消除，并得到K*K分，直到所有方块都消除。显然，不同的消除顺序得分不同，薯队长希望您能告诉他，这个游戏最多能得到多少分。



# 输入描述
# 第一行一个数字N

# 第二个N个数字，为每个方块的颜色

# 输出描述
# 输出只有一行，即最多能得到多少分

# 比如如下例子，先消除颜色2的部分，得到16分，接下来消除颜色1的部分，得到9分，最后消除颜色3的部分，得到4分，一共得到29分


# 样例输入
# 9
# 3 2 2 2 2 1 1 1 3
# 样例输出
# 29
n = int(input())
nums = input().split(" ")
for i in range(n):
    nums[i] = int(nums[i])

global max_bonus
max_bonus = 0
# 遍历每一个起点，做消除，注意数组的引用
def dfs(array, bonus):
    length = len(array)
    if length == 0:
        global max_bonus
        max_bonus = max(max_bonus, bonus)
        return
    for i in range(length):
        if i > 0 and array[i] == array[i - 1]:
            continue
        left, right, curr_length = i - 1, i + 1, 1
        while (left >= 0 and array[left] == array[i]) or (right < length and array[right] == array[i]):
            if left >= 0 and array[left] == array[i]:
                left -= 1
                curr_length += 1
            if right < length and array[right] == array[i]:
                curr_length += 1
                right += 1
        # 数组出现变化
        dfs(array[:left+1] + array[right:], bonus + curr_length * curr_length)

dfs(nums, 0)
print(max_bonus)