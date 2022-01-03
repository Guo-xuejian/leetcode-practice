#
# @lc app=leetcode.cn id=649 lang=python3
#
# [649] Dota2 参议院
#

# @lc code=start
class Solution:
    def predictPartyVictory(self, senate: str) -> str:
        n = len(senate)
        radiant = collections.deque()   # radiant 队列
        dire = collections.deque()      # dire 队列

        for i, ch in enumerate(senate):  # 记录两方议员的索引
            if ch == "R":   # radiant 入队
                radiant.append(i)
            else:   # dire 入队
                dire.append(i)

        while radiant and dire:
            # 队首比较，如果位次靠前则该议员禁止对方议员参与后续所有投票
            if radiant[0] < dire[0]:    # radiant 禁止 dire
                radiant.append(radiant[0] + n)  # 同时改议员进入队尾等待下次他的投票权的来临
            else:       # dire 禁止 radiant
                dire.append(dire[0] + n)        # 同上
            radiant.popleft()   # 不论如何出队，被禁止则永久出队，未被禁止则去到队尾
            dire.popleft()
        # 最终只会剩下一方议员，判定哪一方不为空即可
        return "Radiant" if radiant else "Dire"

# @lc code=end
