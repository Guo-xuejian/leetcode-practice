#
# @lc app=leetcode.cn id=1404 lang=python3
#
# [1404] 将二进制表示减到 1 的步骤数
#

# @lc code=start
class Solution:
    def numSteps(self, s: str) -> int:
        # 模拟法
        # num = int(s, 2)
        # moves = 0
        # while num != 1:
        #     if num % 2 == 0:
        #         num >>= 1
        #     else:
        #         num += 1
        #     moves += 1
        # return moves
        
        # 逐位判定
        index = len(s) - 1
        one_status = False
        res = 0
        while index > 0:    # 等于情况不行，最终为 1 时不能继续判定
            # 当前位为 1 或者当前为应该从 0 变成 1
            if s[index] == "1" or one_status:
                res += 2        # 奇数加 1 除 2，两步
                index -= 1

                while index >= 0 and s[index] == "1":
                    # 连接着的所有 1 都变成 0，偶数情况除 2，只需一步
                    res += 1
                    index -= 1
                
                if index >= 0:
                    # 退出上一个循环时当前 index 对应为 0，需要变成 1，oneStatus 记录
                    one_status = True
            else:   # 为 0 ，偶数
                res += 1
                index -= 1
        return res

# @lc code=end

