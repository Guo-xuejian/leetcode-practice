# 五进制数组转整数
# 详细描述
# 给定一个整形数组，此数组是一个整数数字的五进制表示形式，数组中每个元素的值都小于等于4
# 请你返回该数组所表示数字的 十进制值 。
# 注意：
# 1.数组不为空
# 2.数组的节点数不超过20
# 其他
# 时间限制: 1000ms
# 内存限制:256.0MB

# 输入输出示例
# 示例1
# 输入
# [4,3,1]
# 输出
# 116

# Note: 类名、方法名、参数名已经指定，请勿修改
# 将一个以数组表示的五进制数，转换成十进制整数
# @param nums int整型 一维数组 五进制数数组
# @return long长整型
#
class Solution:
    def Convert(self, nums) :
        # 1 + 3 * 5 + 4 * 25
        m = len(nums)
        multi_num = 0
        res = 0
        for i in range(m - 1, -1, -1):
            res += nums[i] * (5**(multi_num))
            multi_num += 1
        return res


# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++


# IPO
# 详细描述
# 假设 A公司 即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，A公司 希望在 IPO 之前开展一些项目以增加其资本。 由于资源有限，它只能在 IPO 之前完成最多 k 个不同的项目。现需要你帮助 A公司 设计完成最多 k 个不同项目后得到最大总资本的方式。
# 给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。
# 最初，你的资本为 w 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。
# 总而言之，从给定项目中选择 最多 k 个不同项目的列表，以 最大化最终资本 ，并输出最终可获得的最多资本。
# 备注： 答案保证在 32 位有符号整数范围内。
# 其他
# 时间限制: 1000ms
# 内存限制:256.0MB

# 输入输出示例
# 示例1
# 输入
# 2,0,[1,2,3],[0,1,1]
# 输出

# 4
# 说明
# 由于你的初始资本为 0，你仅可以从 0 号项目开始。
# 在完成后，你将获得 1 的利润，你的总资本将变为 1。
# 此时你可以选择开始 1 号或 2 号项目。
# 由于你最多可以选择两个项目，所以你需要完成 2 号项目以获得最大的资本。
# 因此，输出最后最大化的资本，为 0 + 1 + 3 = 4。
# 示例2
# 输入
# 2,0,[1,2,3],[1,1,1]
# 输出

# 0
# 说明
# 由于你的初始资本为 0，不满足启动任何一个项目的最小资本，故最大化的资本仍为0
# 示例3
# 输入
# 2,1,[1,2,3],[1,1,1]
# 输出

# 6
# 说明
# 由于你的初始资本为 1，三个项目均可以参与投资，故首次选择最大利润的2号项目进行投资
# 在完成后，你将获得 3 的利润，你的总资本将变为 4。
# 此时你可以选择开始 0 号或 1 号项目。
# 由于你最多可以选择两个项目，所以你需要完成 1 号项目以获得最大的资本。
# 因此，输出最后最大化的资本，为 1 + 3 + 2 = 6



#
# Note: 类名、方法名、参数名已经指定，请勿修改
# @param k int整型  
# @param w int整型  
# @param profits int整型 一维数组 
# @param capital int整型 一维数组 
# @return int整型
#
class Solution:
    def findMaximizedCapital(self, k, w, profits, capital) :
        if k == 0:
            return w
        # 字典加速访问
        # profit_cost = {capital[i]: profits[i] for i in range(len(profits))}
        profit_cost = {}
        for i in range(len(profits)):
            if capital[i] not in profit_cost:
                profit_cost[capital[i]] = [profits[i],]
            else:
                (profit_cost[capital[i]]).append(profits[i])
        # 先模拟一次的情况
        # if w in profit_cost:
        # 当前成本可以完成至少一个项目
        # curr_money = w
        while k > 0:
            max_money = w
            max_key = None
            max_value = None
            for key in profit_cost.keys():
                if key <= w:   # 当前资本可以完成该项目
                    curr_max = max(profit_cost[key])
                    if w + curr_max > max_money:
                        max_key = key
                        max_value = curr_max
                        max_money = w + curr_max
            # 一轮循环结束，得到当前所能做得项目获得收益
            if max_key is None:  # 如果某一轮中无法完成做项目，则后续也不能，退出外层循环
                break
            w = max_money
            (profit_cost[max_key]).remove(max_value)    # 项目完成后续不能再做
            if len(profit_cost[max_key]) == 0:  # 为空时该成本项目已已经做完
                del profit_cost[max_key]
            k -= 1
        return w


# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++


# 链表合并
# 详细描述
# 将两个 降序 链表合并为一个新的 升序 链表并返回，新链表是通过拼接给定的两个链表的所有节点组成的
# 其他
# 时间限制: 1000ms
# 内存限制:256.0MB

# 输入输出示例
# 示例1
# 输入
# {4,3,2,1},{9,8,7,6}
# 输出
# {1,2,3,4,6,7,8,9}


# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

# Note: 类名、方法名、参数名已经指定，请勿修改
# 合并两个降序链表为一个升序链表
# @param l1 ListNode类  降序链表
# @param l2 ListNode类  降序链表
# @return ListNode类
#
class Solution:
    def MergeList(self, l1, l2) :
        dummy_node = ListNode(0)
        while l1 or l2:
            new_node = ListNode(0)
            if l1 and l2:
                # 两者都未空
                if l1.val > l2.val:
                    new_node.val = l1.val
                    l1 = l1.next
                else:
                    new_node.val = l2.val
                    l2 = l2.next
            elif l1:
                # l2 已空
                new_node = ListNode(l1.val)
                l1 = l1.next
            else:
                # l1 已空
                new_node = ListNode(l2.val)
                l2 = l2.next
            new_node.next = dummy_node.next
            dummy_node.next = new_node
        return dummy_node.next
