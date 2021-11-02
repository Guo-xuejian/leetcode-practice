#
# @lc app=leetcode.cn id=22 lang=python3
#
# [22] 括号生成
#

# @lc code=start
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        def generate(A):
            if len(A) == 2*n:       # 长度要符合要求
                if valid(A):        # 也得是有效括号
                    ans.append("".join(A))
            else:
                # 多叉树，每一条分支都会走到长度符合要求的情况才会结束
                A.append('(')
                generate(A)
                A.pop()
                A.append(')')
                generate(A)
                A.pop()

        def valid(A):
            bal = 0     # >0 不符合要求  <0 也不会符合要求，只有正好才符合要求
            for c in A:
                if c == '(': bal += 1
                else: bal -= 1
                if bal < 0: return False
            return bal == 0

        ans = []
        generate([])
        return ans

# @lc code=end

