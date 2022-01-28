#
# @lc app=leetcode.cn id=1769 lang=python3
#
# [1769] 移动所有球到每个盒子所需的最小操作数
#

# @lc code=start
class Solution:
    def minOperations(self, boxes: str) -> List[int]:
        length = len(boxes)

        def find_boxes(index):
            res = 0
            left, right = index - 1, index + 1
            
            while left >= 0 or right <= length - 1:
                if left >= 0:
                    if boxes[left] == "1":
                        res += index - left
                    left -= 1
                if right <= length - 1:
                    if boxes[right] == "1":
                        res += right - index
                    right += 1
            return res

        res = []
        for i in range(length):
            res.append(find_boxes(i))
        return res
# @lc code=end

