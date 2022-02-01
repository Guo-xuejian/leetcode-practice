#
# @lc app=leetcode.cn id=125 lang=python3
#
# [125] 验证回文串
#

# @lc code=start
class Solution:
    def isPalindrome(self, s: str) -> bool:
        left, right = 0, len(s) - 1

        def validPart(one):     # 字符类型判定
            if one.isdigit() or ord("a") <= ord(one) <= ord("z") or ord("A") <= ord(one) <= ord("Z"):
                return True
            return False

        while left < right:
            num_status = False
            while left < right and not validPart(s[left]):
                left += 1
            while left < right and not validPart(s[right]):
                right -= 1
            if s[left].isdigit() or s[right].isdigit():     # 存在数字
                num_status = True
            if s[left] == s[right] or ((ord(s[left]) + 32 == ord(s[right]) or ord(s[right]) + 32 == ord(s[left])) and not num_status):     # 字符相同继续
                left += 1
                right -= 1
            else:   # 不相同退出
                return False
        # 无不同
        return True

# @lc code=end
