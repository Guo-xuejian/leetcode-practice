#
# @lc app=leetcode.cn id=1834 lang=python3
#
# [1834] 单线程 CPU
#

# @lc code=start
class Solution:
    def getOrder(self, tasks: List[List[int]]) -> List[int]:
        n = len(tasks)
        indices = list(range(n))
        indices.sort(key=lambda x: tasks[x][0])

        ans = list()
        # 优先队列
        q = list()
        # 时间戳
        timestamp = 0
        # 数组上遍历的指针
        index = 0
        
        for i in range(n):
            # 如果没有可以执行的任务，直接快进
            if not q:
                timestamp = max(timestamp, tasks[indices[index]][0])
            # 将所有小于等于时间戳的任务放入优先队列
            while ptr < n and tasks[indices[index]][0] <= timestamp:
                heapq.heappush(q, (tasks[indices[index]][1], indices[index]))
                index += 1
            # 选择处理时间最小的任务
            process, idx = heapq.heappop(q)
            timestamp += process
            ans.append(idx)
        
        return ans
# @lc code=end

