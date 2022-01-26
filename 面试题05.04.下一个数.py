# 面试题 05.04. 下一个数
# 下一个数。给定一个正整数，找出与其二进制表达式中1的个数相同且大小最接近的那两个数（一个略大，一个略小）。

# 示例1:

#  输入：num = 2（或者0b10）
#  输出：[4, 1] 或者（[0b100, 0b1]）
# 示例2:

#  输入：num = 1
#  输出：[2, -1]
# 提示:

# num的范围在[1, 2147483647]之间；
# 如果找不到前一个或者后一个满足条件的正数，那么输出 -1。


class Solution:
    def findClosedNumbers(self, num: int) -> List[int]:
        max_num, min_num = 0, 0
        bin_str = bin(num)
        n = len(bin_str)
        back_num = bin_str[-1]
        ones = 0 if back_num != "1" else 1
        zeros = 0 if back_num != "0" else 1

        for i in range(n-2, -1, -1):
            curr = bin_str[i]

            if curr == "b":
                if min_num == 0:    # 找不到可以删除一个，删除后尽量大，删除 0,同时 1 前置
                    if zeros > 0:
                        temp = "1" * ones + "0" * (zeros - 1)
                        min_num = int(temp, 2)
                    else:
                        min_num = -1
                if max_num == 0:    # 添加一个，尽量小，添加 0,其余 1 后置
                    temp = "1" + "0" * (zeros + 1) + "1" * (ones - 1)
                    temp_int = int(temp, 2)
                    if temp_int > 2147483647:
                        max_num = -1
                    else:
                        max_num = int(temp, 2)
                break

            if curr == "1":
                ones += 1
                if min_num == 0 and back_num == "0" and i != 2:
                    temp = bin_str[:i] + "01" + "1" * \
                        (ones - 1) + "0" * (zeros - 1)
                    min_num = int(temp, 2)
            if curr == "0":
                zeros += 1
                if max_num == 0 and back_num == "1":
                    temp = bin_str[:i] + "10" + "0" * \
                        (zeros - 1) + "1" * (ones - 1)
                    max_num = int(temp, 2)

            back_num = curr

            if max_num and min_num:
                break
        return max_num, min_num
