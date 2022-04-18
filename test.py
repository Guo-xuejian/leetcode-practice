# 求最短子串
# 详细描述
# 给定两个字符串s和t，在s中找出包括t中所有字符的第一个最小字符串。如果没有则返回空字符串。

 

# 说明：字符串仅包含英文字符。长度不限且假设单机能正常处理。

# 其他
# 时间限制: 1000ms

# 内存限制:256.0MB

# 输入输出示例
# 示例1
# 输入
# 复制
# "ADAOBECODEBANCB","ABC"
# 输出
# 复制
# "BANC"



#
# Note: 类名、方法名、参数名已经指定，请勿修改
#
#
# 
# @param s string字符串  
# @param t string字符串  
# @return string字符串
#
class Solution:
    def find_min_str(self, s, t) :
        # write code here
        m, n = len(s), len(t)
        if m < n:
            return ""
        letter_dict = {}
        for letter in t:
            letter_dict[letter] = 1
        res = ""
        small_count = float("inf")
        for i in range(m - n):
            # if i > 0 and s[i] == s[i - 1]:
            #     continue
            if s[i] in letter_dict:
                for j in range(i + n, m):
                    curr = s[i:j]
                    status = True
                    for letter in t:
                        if letter not in curr:
                            status = False
                            break
                    if status and small_count > j - i:
                        small_count = j - i
                        res = curr
                        break
        return res
                    



# 2.
# 编程题/核心代码
# (20分)
# 本题无跳出页面限制，支持使用本地IDE调试，最终以系统调试结果为准。
# 输入一个整数，输出该整数二进制表示中1的个数
# 详细描述
# 输入一个整数，输出该整数二进制表示中1的个数

# 要包含负数处理过程

 

# 在计算机中，负数以其正值的补码形式表达
# 什么叫补码呢？这得从原码，反码说起。
# 原码：一个整数，按照绝对值大小转换成的二进制数，称为原码。
# 比如 00000000 00000000 00000000 00000101 是 5的 原码。
# 反码：将二进制数按位取反，所得的新二进制数称为原二进制数的反码。
# 取反操作指：原为1，得0；原为0，得1。（1变0; 0变1）
# 比如：将00000000 00000000 00000000 00000101每一位取反，得11111111 11111111 11111111 11111010。
# 称：11111111 11111111 11111111 11111010是 00000000 00000000 00000000 00000101 的反码。
# 反码是相互的，所以也可称：
# 11111111 11111111 11111111 11111010 和00000000 00000000 00000000 00000101 互为反码。
# 补码：反码加1称为补码。
# 也就是说，要得到一个数的补码，先得到反码，然后将反码加上1，所得数称为补码。
# 比如：00000000 00000000 00000000 00000101的反码是：11111111 11111111 11111111 11111010。
# 那么，补码为：
# 11111111 11111111 11111111 11111010 + 1 = 11111111 11111111 11111111 11111011
# 所以，-5 在计算机中表达为：11111111 11111111 11111111 11111011

# 其他
# 时间限制: 1000ms

# 内存限制:256.0MB

# 输入输出示例
# 示例1
# 输入
# 复制
# 9
# 输出
# 复制
# 2
# 示例2
# 输入
# 复制
# 1023
# 输出
# 复制
# 10
# 示例3
# 输入
# 复制
# 0
# 输出
# 复制
# 0
# 示例4
# 输入
# 复制
# -5
# 输出
# 复制
# 31




#
# Note: 类名、方法名、参数名已经指定，请勿修改
#
#
# 
# @param input_int int整型  
# @return int整型
#
class Solution:
    def oct_to_binary(self, input_int) :
        # write code here
        res = 0
        if input_int < 0:
            one = 1
            maxInt = 0
            for i in range(32):
                maxInt += one
                one <<= 1
            input_int *= -1
            input_int ^= maxInt
            input_int += 1
        while input_int > 0:
            if input_int & 1 == 1:
                res += 1
            input_int >>= 1
        return res