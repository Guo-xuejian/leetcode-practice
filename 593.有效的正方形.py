#
# @lc app=leetcode.cn id=593 lang=python3
#
# [593] 有效的正方形
#
# https://leetcode.cn/problems/valid-square/description/
#
# algorithms
# Medium (44.10%)
# Likes:    156
# Dislikes: 0
# Total Accepted:    31.1K
# Total Submissions: 66.1K
# Testcase Example:  '[0,0]\n[1,1]\n[1,0]\n[0,1]'
#
# 给定2D空间中四个点的坐标 p1, p2, p3 和 p4，如果这四个点构成一个正方形，则返回 true 。
# 
# 点的坐标 pi 表示为 [xi, yi] 。输入 不是 按任何顺序给出的。
# 
# 一个 有效的正方形 有四条等边和四个等角(90度角)。
# 
# 
# 
# 示例 1:
# 
# 
# 输入: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
# 输出: True
# 
# 
# 示例 2:
# 
# 
# 输入：p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,12]
# 输出：false
# 
# 
# 示例 3:
# 
# 
# 输入：p1 = [1,0], p2 = [-1,0], p3 = [0,1], p4 = [0,-1]
# 输出：true
# 
# 
# 
# 
# 提示:
# 
# 
# p1.length == p2.length == p3.length == p4.length == 2
# -10^4 <= xi, yi <= 10^4
# 
# 
#

# @lc code=start
def checkLength(v1: Tuple[int, int], v2: Tuple[int, int]) -> bool:
    return v1[0] * v1[0] + v1[1] * v1[1] == v2[0] * v2[0] + v2[1] * v2[1]

def checkMidPoint(p1: List[int], p2: List[int], p3: List[int], p4: List[int]) -> bool:
    return p1[0] + p2[0] == p3[0] + p4[0] and p1[1] + p2[1] == p3[1] + p4[1]

def calCos(v1: Tuple[int, int], v2: Tuple[int, int]) -> int:
    return v1[0] * v2[0] + v1[1] * v2[1]

def help(p1: List[int], p2: List[int], p3: List[int], p4: List[int]) -> bool:
    v1 = (p1[0] - p2[0], p1[1] - p2[1])
    v2 = (p3[0] - p4[0], p3[1] - p4[1])
    return checkMidPoint(p1, p2, p3, p4) and checkLength(v1, v2) and calCos(v1, v2) == 0

class Solution:
    def validSquare(self, p1: List[int], p2: List[int], p3: List[int], p4: List[int]) -> bool:
        if p1 == p2:
            return False
        if help(p1, p2, p3, p4):
            return True
        if p1 == p3:
            return False
        if help(p1, p3, p2, p4):
            return True
        if p1 == p4:
            return False
        if help(p1, p4, p2, p3):
            return True
        return False
# @lc code=end

