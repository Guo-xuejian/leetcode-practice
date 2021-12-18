#
# @lc app=leetcode.cn id=1333 lang=python3
#
# [1333] 餐厅过滤器
#

# @lc code=start
class Solution:
    def filterRestaurants(self, restaurants: List[List[int]], veganFriendly: int, maxPrice: int, maxDistance: int) -> List[int]:
        nums = []
        for restaurant in restaurants:
            if veganFriendly == (veganFriendly & restaurant[2]) and restaurant[3] <= maxPrice and restaurant[4] <= maxDistance:
                nums.append([restaurant[0], restaurant[1]])
        nums.sort(key = lambda x: (x[1], x[0]), reverse = True)
        return [x[0] for x in nums]
# @lc code=end

