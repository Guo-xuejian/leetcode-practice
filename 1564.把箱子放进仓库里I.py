# 1564. 把箱子放进仓库里 I
# 给定两个正整数数组 boxes 和 warehouse ，分别包含单位宽度的箱子的高度，以及仓库中 n 个房间各自的高度。仓库的房间分别从 0 到 n - 1 自左向右编号， warehouse[i] （索引从 0 开始）是第 i 个房间的高度。

# 箱子放进仓库时遵循下列规则：

# 箱子不可叠放。
# 你可以重新调整箱子的顺序。
# 箱子只能从左向右推进仓库中。
# 如果仓库中某房间的高度小于某箱子的高度，则这个箱子和之后的箱子都会停在这个房间的前面。
# 你最多可以在仓库中放进多少个箱子？

 

# 示例 1：



# 输入：boxes = [4,3,4,1], warehouse = [5,3,3,4,1]
# 输出：3
# 解释：

# 我们可以先把高度为 1 的箱子放入 4 号房间，然后再把高度为 3 的箱子放入 1 号、 2 号或 3 号房间，最后再把高度为 4 的箱子放入 0 号房间。
# 我们不可能把所有 4 个箱子全部放进仓库里。
# 示例 2：



# 输入：boxes = [1,2,2,3,4], warehouse = [3,4,1,2]
# 输出：3
# 解释：

# 我们注意到，不可能把高度为 4 的箱子放入仓库中，因为它不能通过高度为 3 的房间。
# 而且，对于最后两个房间 2 号和 3 号来说，只有高度为 1 的箱子可以放进去。
# 我们最多可以放进 3 个箱子，如上图所示。黄色的箱子也可以放入 2 号房间。
# 交换橙色和绿色箱子的位置，或是将这两个箱子与红色箱子交换位置，也是可以的。
# 示例 3：

# 输入：boxes = [1,2,3], warehouse = [1,2,3,4]
# 输出：1
# 解释：由于第一个房间的高度为 1，我们只能放进高度为 1 的箱子。
# 示例 4：

# 输入：boxes = [4,5,6], warehouse = [3,3,3,3,3]
# 输出：0
 

# 提示：

# n == warehouse.length
# 1 <= boxes.length, warehouse.length <= 10^5
# 1 <= boxes[i], warehouse[i] <= 10^9


class Solution:
    def maxBoxesInWarehouse(self, boxes: List[int], warehouse: List[int]) -> int:
        length = len(warehouse)
        for i in range(1, length):
            # 题目限制，所有后面的房间中不会出现比前面房间中更高的箱子
            warehouse[i] = min(warehouse[i], warehouse[i-1])
        boxes.sort(reverse = True)
        index = len(boxes) - 1
        num = 0
        for i in range(length - 1, -1, -1):
            # boxes 从小到大 从后往前，warehouse 先匹配小的
            if index >= 0 and boxes[index] <= warehouse[i]:
                num += 1
                index -= 1
        return num