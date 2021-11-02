#
# @lc app=leetcode.cn id=17 lang=python3
#
# [17] 电话号码的字母组合
#

# @lc code=start
class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        # 字典存储信息
        num_lettere_dict = {
            "2":"abc",
            "3":"def",
            "4":"ghi",
            "5":"jkl",
            "6":"mno",
            "7":"pqrs",
            "8":"tuv",
            "9":"wxyz",
        }
        if not digits:
            return []
        def track(index:int):
            if index == lengthOfDigits:
                result_list.append("".join(temp))
            else:
                digit = digits[index]
                for letter in num_lettere_dict[digit]:
                    temp.append(letter)
                    track(index + 1)
                    temp.pop()
        lengthOfDigits = len(digits)
        result_list = list()
        temp = list()
        track(0)
        return result_list
# @lc code=end

