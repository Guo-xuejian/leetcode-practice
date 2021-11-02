#
# @lc app=leetcode.cn id=20 lang=python3
#
# [20] 有效的括号
#

# @lc code=start
class Solution:
    def isValid(self, s: str) -> bool:
        # 数量上不符合要求直接返回 false
        if len(s) % 2 == 1:
            return False
        
        pairs = {
            ")": "(",
            "]": "[",
            "}": "{"
        }
        # 列表模拟栈
        stack = list()
        for one in s:
            # 只处理右括号，右括号才是需要对栈的结构产生缩减的起点
            if one in pairs:
                # 如果栈为空或者最后一个元素不是该右括号对应的左括号，那么明显不符合条件
                if not stack or stack[-1] != pairs[one]:
                    return False
                # 符合条件，栈中出栈对应的配对左括号
                stack.pop()
            else:
                # 左括号入栈
                stack.append(one)
        # 栈为空则全部匹配，栈中仍有元素意味着不符合条件
        return not stack
# @lc code=end

