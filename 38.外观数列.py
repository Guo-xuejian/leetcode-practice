#
# @lc app=leetcode.cn id=38 lang=python3
#
# [38] 外观数列
#

# @lc code=start
class Solution:
    def countAndSay(self, n: int) -> str:
        prev = "1"
        for i in range(n-1):
            curr = ""
            pos = 0
            start = 0
            while pos < len(prev):
                while pos < len(prev) and prev[pos] == prev[start]:
                    # 相同数字 pos 计数
                    pos += 1
                # 不同数字中断并给出一部分结果
                curr += str(pos - start) + prev[start]
                # 设置新的起点，继续
                start = pos
            # 一轮遍历结束覆盖 prev，作为下一轮遍历的字符串
            prev = curr
        return prev
# @lc code=end

