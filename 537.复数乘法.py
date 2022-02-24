#
# @lc app=leetcode.cn id=537 lang=python3
#
# [537] 复数乘法
#

# @lc code=start
class Solution:
    def complexNumberMultiply(self, num1: str, num2: str) -> str:
        # 拆分字符串，因为有共同点，都只有一个 + 号
        split1, split2 = num1.split("+"), num2.split("+")
        pre1, late1 = int(split1[0]), int(split1[1].replace("i", ""))
        pre2, late2 = int(split2[0]), int(split2[1].replace("i", ""))
        final_pre, final_late = pre1 * pre2 - (late1 * late2), pre1 * late2 + pre2 * late1
        return f"{final_pre}+{final_late}i"
# @lc code=end

