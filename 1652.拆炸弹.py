#
# @lc app=leetcode.cn id=1652 lang=python3
#
# [1652] 拆炸弹
#

# @lc code=start
class Solution:
    def decrypt(self, code: List[int], k: int) -> List[int]:
        code_length = len(code)
        if k == 0:  # 为 0 时按要求返回全零值
            return [0]*code_length
        if k == 1:  # k==1 时只需要后移一位即可，但是首位需要接到最后
            code.append(code[0])
            return code[1:]
        final_list = []
        if k > 0:
            if k == code_length - 1:    # 长度为此时恰好计算总和减去对应位置数字即可
                sum_all = sum(code)
                for i in range(code_length):
                    code[i] = sum_all - code[i]
                return code
            for i in range(code_length):    # 其余情况
                curr_num = 0
                index = i + 1 if i + 1 < code_length else 0   # 索引前移
                final_index = i + k if i + \
                    k < code_length else (i + k) % code_length   # 最终索引不能越界
                while index != final_index:
                    curr_num += code[index]
                    index = index + 1 if index + 1 < code_length else 0
                curr_num += code[index]  # 补齐最后一个数字
                final_list.append(curr_num)
        elif k < 0:
            for i in range(code_length):
                curr_num = 0
                index = i - 1 if i - 1 >= 0 else code_length + i - 1    # 越界判定
                final_index = i + k if i + k >= 0 else code_length + i + k
                while index != final_index:
                    curr_num += code[index]
                    index = index - 1 if index - 1 >= 0 else code_length - 1
                curr_num += code[index]  # 补齐最后一个数字
                final_list.append(curr_num)
        return final_list
# @lc code=end
