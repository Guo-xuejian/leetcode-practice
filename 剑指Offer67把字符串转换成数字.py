class Solution:
    def strToInt(self, str: str) -> int:
        str = str.strip()
        temp = ''
        flag = ''
        res = 0
        for i in range(len(str)):
            if i == 0 and str[i] != '-' and str[i] != '+':
                if not str[i].isdigit():
                    break
            if i == 0 and str[i] == '-' or  i == 0 and str[i] == '+':
                flag += str[i]
            elif str[i].isdigit():
                if str[i] == '.':
                    isDot = True
                temp += str[i]
            else:
                break
        if temp == '':
            return 0
        res =  int(flag + temp)
        if res < -2**31:
            res = -2147483648
        elif res >= 2**31:
            res = (2**31-1)
        return res