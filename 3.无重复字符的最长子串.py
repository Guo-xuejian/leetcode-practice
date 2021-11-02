#
# @lc app=leetcode.cn id=3 lang=python3
#
# [3] 无重复字符的最长子串
#

# @lc code=start
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        length = len(s)
        lengthOfUniqe_list = len(list(set(s)))
        if lengthOfUniqe_list == length:
            return length
        # a 记录一个正在检查是否符合要求的子串元素列表，只要哦后一个元素不存在于 a 中，那么就会加入 a
        # lengthOfLongestStr 会在遍历元素已经存在于 a 中时生效，将 a 中元素整体数量和 b 比较后取较大值，a 截断至此次遍历元素在 a 未截断之前索引之后，持续保证 a 中元素单一
        a = []
        lengthOfLongestStr = 0
        for i in range(length):
            if s[i] not in a:
                # 不存在于 a 便加入
                a.append(s[i])
            else:
                # 存在于 a 则计算 a 中元素数量并进行比较
                lengthOfLongestStr = max(lengthOfLongestStr, len(a))
                # 得出 a 中 s[i] 的索引值并对 a 进行截断
                last_element_index = a.index(s[i])
                a = a[last_element_index+1:]
                # 之后加入本次元素
                a.append(s[i])
        # 最后还需要一次判定用来取得最终的最大值并返回结果
        return max(lengthOfLongestStr, len(a))
        
# @lc code=end

