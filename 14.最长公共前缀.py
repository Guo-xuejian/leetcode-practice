#
# @lc app=leetcode.cn id=14 lang=python3
#
# [14] 最长公共前缀
#

# @lc code=start
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""
        prefix = strs[0]
        count = len(strs)
        for i in range(count):
            prefix = self.findCommonPart(prefix, strs[i])
            # 当 前缀长度为 0 时跳出，因为所有的字符串的公共部分都是空字符串，之后的元素也不需要再进行判定
            if len(prefix) == 0:
                break
        return prefix
    
    def findCommonPart(self, str1, str2):
        # 首先需要对比两个字符串长度，取较小的，避免数组越界
	    # 从第一个元素开始，因为可能存在前面有共同部分，后面元素没有共同部分
        lengthOfMinPart = min(len(str1), len(str2))
        index = 0
        while index < lengthOfMinPart and str1[index] == str2[index]:
            index += 1
        return str1[:index]
        
        
# @lc code=end

