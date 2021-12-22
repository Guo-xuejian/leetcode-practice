class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if not s:
            return 0
        if len(s) == 1:
            return 1 
        long_str = ''
        length = 0
        final_length = 0
        for one in s:
            if one in long_str:
                one_index = long_str.index(one)
                # 获取long_str中该字母的下标并切分改字符串
                long_str = long_str[one_index:].replace(one, '')
                length -= one_index + 1
                # 将该字符拼接上去，之后的加一不再上面是为了更清晰作用，one_index不一定会是0，该字符如果处在long_str的中间，那么这么处理才是合适的
                long_str += one
                length += 1
                continue
            # 如果long_str中没有该字母
            long_str += one
            length += 1
            if length > final_length:
                final_length = length
        return final_length