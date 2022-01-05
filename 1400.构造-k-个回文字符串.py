#
# @lc app=leetcode.cn id=1400 lang=python3
#
# [1400] 构造 K 个回文字符串
#

# @lc code=start
class Solution:
    def canConstruct(self, s: str, k: int) -> bool:
        # 单数（不仅包括单个，三个也会存在类似问题）字符的数量不能大于 k
        # 同时出现次数大于 2 的字符不能少于 k !!!!!
        # 上方错误，可以少于 k，单个字符也是回文
        s_length = len(s)
        if s_length == k:   # 长度相等，则全部拆成单个字符，单个字符也是回文
            return True
        elif s_length < k:  # 长度小于则必然不可能拆成 k 个字符串
            return False
        letter_num_dict = {}    # 字典记录每个字符出现的次数
        for letter in s:
            if letter in letter_num_dict:
                letter_num_dict[letter] += 1
            else:
                letter_num_dict[letter] = 1
        single_letter = 0   # 单数字符个数
        # not_single_letter = 0   # 出现超 2 的字符数量
        for val in letter_num_dict.values():
            if val % 2 != 0:    # 单数字符
                single_letter += 1
                # if val == 1:    # 为 1 时退出本次循环
                #     continue
            # 大于 1
            # not_single_letter += 1
        if single_letter > k:   # 如果单数字符大于 k ，则无法通过分配单个字符为回文串与单个字符与双数字符组合方式得到 k 个字符串
            return False
        return True
# @lc code=end
