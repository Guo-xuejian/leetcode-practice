#
# @lc app=leetcode.cn id=1626 lang=python3
#
# [1626] 无矛盾的最佳球队
#

# @lc code=start
class Solution:
    def bestTeamScore(self, scores: List[int], ages: List[int]) -> int:
        age_sorted_scores = [score for _, score in sorted(zip(ages, scores))]
        age_sorted_scores = [0] + age_sorted_scores # 哨兵，简化，持续和之前的所有年龄小于等于者以及分数小于者做序列总和比较
        dp = [0] * len(age_sorted_scores)
        for i in range(1, len(age_sorted_scores)):
            for j in range(i):
                if age_sorted_scores[i] >= age_sorted_scores[j]:
                    dp[i] = max(dp[i], dp[j] + age_sorted_scores[i]) # 取以 i 位置的最大和
        return max(dp)

# @lc code=end

