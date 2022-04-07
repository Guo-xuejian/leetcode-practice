#
# @lc app=leetcode.cn id=703 lang=python3
#
# [703] 数据流中的第 K 大元素
#

# @lc code=start
import heapq
class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        self._minHeap = []
        heapq.heapify(self._minHeap)
        self.k = k
        self.nums = nums
        if len(self.nums) > self.k:
            for i in range(self.k):
                heapq.heappush(self._minHeap, self.nums[i])
            for i in range(self.k, len(self.nums)):
                if nums[i] > self._minHeap[0]:
                    heapq.heappop(self._minHeap)
                    heapq.heappush(self._minHeap, self.nums[i])
        else:
            for i in range(len(self.nums)):
                heapq.heappush(self._minHeap, self.nums[i])

    def add(self, val: int) -> int:
        self.nums.append(val)
        if len(self.nums) <= self.k:
            heapq.heappush(self._minHeap, val)
        else:
            if val > self._minHeap[0]:
                heapq.heappop(self._minHeap)
                heapq.heappush(self._minHeap, val)
        return self._minHeap[0]


# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val)
# @lc code=end

