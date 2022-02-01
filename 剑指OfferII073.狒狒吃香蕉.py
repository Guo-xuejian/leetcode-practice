# 剑指 Offer II 073. 狒狒吃香蕉
# 狒狒喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。

# 狒狒可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉，下一个小时才会开始吃另一堆的香蕉。

# 狒狒喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

# 返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。


# 示例 1：

# 输入: piles = [3,6,7,11], H = 8
# 输出: 4
# 示例 2：

# 输入: piles = [30,11,23,4,20], H = 5
# 输出: 30
# 示例 3：

# 输入: piles = [30,11,23,4,20], H = 6
# 输出: 23


# 提示：

# 1 <= piles.length <= 10^4
# piles.length <= H <= 10^9
# 1 <= piles[i] <= 10^9


class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        left, right = 1, max(piles)
        while left <= right:
            mid = (left + right) // 2
            time_cost = 0
            for pile in piles:
                time_cost += ceil(pile / mid)
                if time_cost > h:   # 大于时满足条件即可退出
                    break
            if time_cost == h:  # 等于,速度刚好满足,试试减小速度
                right = mid - 1
            elif time_cost > h:  # 大于超时,加大速度
                left = mid + 1
            elif time_cost < h:  # 小于和等于一样
                right = mid - 1
        # 退出循环时,只有满足条件的最小速度
        return left
