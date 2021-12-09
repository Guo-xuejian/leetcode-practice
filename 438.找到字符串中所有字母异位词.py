#
# @lc app=leetcode.cn id=438 lang=python3
#
# [438] 找到字符串中所有字母异位词
#

# @lc code=start
from collections import Counter


class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        p_counter = Counter(p)  # 字符映射
        right = len(p)  # 窗口右边界
        range_num = len(s) - right
        result_list = []    # 结果
        for i in range(range_num+1):    # +1 留下了右边界
            temp = Counter(s[i:right])  # 根据窗口计算映射情况
            if temp == p_counter:
                result_list.append(i)   # 相等则写入起始索引
            right += 1  # 右边界移动
        return result_list
# @lc code=end

