#
# @lc app=leetcode.cn id=9 lang=python3
#
# [9] 回文数
#

# @lc code=start
class Solution:
    def isPalindrome(self, x: int) -> bool:
        # 转换为字符串
        string_x = str(x)
        pre = back = 0
        length_of_int = len(string_x)
        # 采用中心扩散的方法
        # 如果长度为偶数,那么前后指针差距为 1
        # 如果长度为奇数,那么前后指针初始值相等
        if length_of_int % 2 == 0:
            pre = int(length_of_int / 2 -1)
            back = pre + 1
        else:
            pre = back = int(length_of_int / 2)
        while True:
            if string_x[pre] != string_x[back]:
                return False
            pre -= 1
            back += 1
            if pre < 0:     # 由于对称,所以之判定 pre 是否越界
                break
        return True
# @lc code=end

