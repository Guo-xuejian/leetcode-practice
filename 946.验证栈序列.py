#
# @lc app=leetcode.cn id=946 lang=python3
#
# [946] 验证栈序列
#

# @lc code=start
class Solution:
    def validateStackSequences(self, pushed: List[int], popped: List[int]) -> bool:
        stack = []
        pop_index = 0
        for num in pushed:
            # get in stack
            stack.append(num)
            while stack and stack[-1] == popped[pop_index]:
                # a element get out of the stack and move pop_index
                stack.pop(-1)
                pop_index += 1
        # if ok, then the stack is empty
        return not stack
# @lc code=end

