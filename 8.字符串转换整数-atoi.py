#
# @lc app=leetcode.cn id=8 lang=python3
#
# [8] 字符串转换整数 (atoi)
#

# @lc code=start
class Solution:
    def myAtoi(self, s: str) -> int:
        # 初始化数字列表
        number_list = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"]
        # 初始化边界值，下一个字符是否允许为+-符号，下一个字符是否允许为空格
        maxInt, signAllowed, whitespaceAllowed = 2 << 30, True, True
        # 初始化正负数判定符号,因为正数与负数的边界值有差异[−2**31,  2**31 − 1],同时初始化一个 digits 用来存储数字
        sign, digits = 1, []
        for one in s:
            # 是空格并且上一步判定允许这一次遍历出现空格===>结束此次遍历
            if one == " " and whitespaceAllowed:
                continue
            if signAllowed:   # 判定下一次遍历是否允许出现符号
                if one == "+":
                    signAllowed = False       # 下一个不能是符号
                    whitespaceAllowed = False  # 下一个不能是空格
                    continue
                elif one == "-":
                    sign = -1
                    signAllowed = False       # 下一个不能是符号
                    whitespaceAllowed = False  # 下一个不能是空格
                    continue
                # 不需要 else,因为只判定数字,而非算式,实际上这部分代码块只会运行一次
            if one not in number_list:  # 字符不在[0...9]内,说明未遇到数字或者数据字符串结束,结束  for 循环
                break
            # 1.允许的空格出现或无空格
            # 2.未出现+-或已经判定
            # 3.c 处在[0...9]范围内
            whitespaceAllowed, signAllowed = False, False   # 正常数字出现,符号空格出现了或者已经判定跳出一次,一个数字不能出现两次+-
            digits.append(int(one)) # 数字写入
        # 初始化place为进制为个位代表的乘数1,(十位也就是10====>方便理解),int8 很有可能不足以支撑
	    # 初始化 num = 0为最终返回数字的初始 
        place, num = 1, 0
        # 判定无效的 0 并截断
        meaningless_zero_index = -1
        for i in range(len(digits)):
            if digits[i] == 0:
                meaningless_zero_index = i  # 如果开始数字为 0,那么就是无效的,那么记录索引至最后一个
            else:
                # 只要第一个数字不是 0 ,那么就是一个正常的数字,结束遍历
                break
        # 判定是否存在无效 0 并截断
        if meaningless_zero_index > -1:
            digits = digits[meaningless_zero_index+1:]
        # 判定数字的正负状态以决定适用的数值边界
        if sign > 0:
            maxInt -= 1
        digits_count = len(digits)
        # 倒序遍历,倒序的好处就是进制的处理比较方便,从个位开始,逐位向上,也比较好记录 place,
        for i in range(digits_count-1, -1, -1):
            num += digits[i] * place
            place *= 10
            if digits_count-i > 10 or num > maxInt:
                # 超过 32位有符号整数,截断为 int(int64(sign) * rtnMax)
                return sign * maxInt
        # 上面为方便处理都是正数,现在还原真实正负状态
        return num * sign

# @lc code=end
