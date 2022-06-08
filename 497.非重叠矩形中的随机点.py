#
# @lc app=leetcode.cn id=497 lang=python3
#
# [497] 非重叠矩形中的随机点
#

# @lc code=start
class Solution:
	
	def __init__(self, rects: List[List[int]]):
		self.rects = rects
		
	def pick(self) -> List[int]:
		result = []
		count = 0
		for rect in self.rects:
			x1, y1, x2, y2 = rect[0], rect[1], rect[2], rect[3]				
			area = (x2 - x1 + 1) * (y2 - y1 + 1)
			count += area    # 当前点的总数
			if randrange(count) < area: # 水塘抽样判断条件
				result = [randrange(x1, x2+1), randrange(y1,y2+1)]
		return result


# Your Solution object will be instantiated and called as such:
# obj = Solution(rects)
# param_1 = obj.pick()
# @lc code=end

