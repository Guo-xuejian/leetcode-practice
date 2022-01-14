#
# @lc app=leetcode.cn id=373 lang=python3
#
# [373] 查找和最小的K对数字
#

# @lc code=start
class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        m, n = len(nums1), len(nums2)
        ans = []
        pq = [(nums1[i] + nums2[0], i, 0) for i in range(min(k, m))]
        while pq and len(ans) < k:
            _, i, j = heappop(pq)   # 最大堆，首位会在poop后改变为最小值，不是常规列表
            ans.append([nums1[i], nums2[j]])
            if j + 1 < n:
                # 删除一个，符合要求加入一个，仍然会维护最大堆特性
                heappush(pq, (nums1[i] + nums2[j + 1], i, j + 1))
        return ans
# @lc code=end

