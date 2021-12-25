#
# @lc app=leetcode.cn id=1418 lang=python3
#
# [1418] 点菜展示表
#

# @lc code=start
from sortedcontainers import SortedSet


class Solution:
    def displayTable(self, orders: List[List[str]]) -> List[List[str]]:
        foods = SortedSet()
        tables = defaultdict(Counter)
        for name,number,food in orders:
            foods.add(food)
            tables[number][food] += 1
        return [["Table"] + [food for food in foods]] + [[table] + [str(tables[table][food]) for food in foods] for table in sorted(tables.keys(), key=int)]


# @lc code=end

