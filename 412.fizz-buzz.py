#
# @lc app=leetcode.cn id=412 lang=python3
#
# [412] Fizz Buzz
#

# @lc code=start
class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        res = []
        # # 两个变量记录当前数字是否能够被3和5整除
        # three_flag = fiveFlag = 0
        # # 遍历 n ，做出判定
        # for i in range(1, n+1):
        #     three_flag = i % 3
        #     fiveFlag = i % 5

        #     # 最高等级判定 ==》同时为3，5的倍数
        #     if not three_flag and not fiveFlag:
        #         res.append("FizzBuzz")
        #     elif not three_flag:
        #         res.append("Fizz")
        #     elif not fiveFlag:
        #         res.append("Buzz")
        #     else:
        #         res.append(str(i))
        # return res
        for i in range(1, n+1):
            temp = ""
            if i%3==0:
                temp += "Fizz"
            if i%5==0:
                temp += "Buzz"
            if len(temp)==0:
                temp += str(i)
            res.append(temp)
        return res

# @lc code=end
