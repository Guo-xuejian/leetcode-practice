#
# @lc app=leetcode.cn id=1996 lang=python3
#
# [1996] 游戏中弱角色的数量
#

# @lc code=start
class Solution:
    def numberOfWeakCharacters(self, properties: List[List[int]]) -> int:
        # 攻击力从大到小排序，防御力从小到大，右边攻击力天然小于等于左边，只需要判定防御力情况即可
        properties.sort(key=lambda x: (-x[0], x[1]))
        ans = 0
        max_defense = 0
        for _, defense in properties:
            if max_defense > defense:   # 比之前的攻击力小的同时防御力也小，当前角色为弱角色
                ans += 1
            else:
                max_defense = max(max_defense, defense)
        return ans
# @lc code=end

