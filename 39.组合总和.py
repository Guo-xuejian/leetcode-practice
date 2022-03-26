#
# @lc app=leetcode.cn id=39 lang=python3
#
# [39] 组合总和
#

# @lc code=start
class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        num_dict = {}
        candidates.sort()
        for num in candidates:
            num_dict[num] = True
        res = []
        length = len(candidates)
        def backTrack(path, curr, index):
            if curr < 0:
                return
            if curr == 0:
                res.append(path[:])
                return
            while index < length:
                path.append(candidates[index])
                backTrack(path, curr - candidates[index], index)
                path.pop()
                index += 1
        backTrack([], target, 0)
        return res
# @lc code=end

