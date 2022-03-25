class TreeNode:
    """树节点"""

    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def reConstructBinaryTree(self, pre, tin):
        """根据前序遍历和中序遍历来重建一颗二叉树，要求前序遍历和中序遍历当中字符不重复
        Args:
          pre: 前序遍历  list类型或者者str类型
          tin: 中序遍历
        Returns
          head: 一个TreeNode类型的根节点
        """
        return self.rebuild_tree(pre, 0, len(pre)-1, tin, 0, len(tin)-1)

    def rebuild_tree(self, pre, pre_start, pre_end, tin, tin_start, tin_end):
        """递归的进行树的重建"""
        if pre_start > pre_end or tin_start > tin_end:
            return None

        head = TreeNode(pre[pre_start])
        tin_mid = tin.index(pre[pre_start])
        left_length = tin_mid - tin_start

        head.left = self.rebuild_tree(
            pre, pre_start+1, pre_start+left_length, tin, tin_start, tin_mid-1)
        head.right = self.rebuild_tree(
            pre, pre_start+left_length+1, pre_end, tin, tin_mid+1, tin_end)

        return head


def post_order_print(head):
    """以后序遍历的方式打印一颗二叉树"""
    if head is None:
        return
    post_order_print(head.left)
    post_order_print(head.right)
    print(head.val, end='')


if __name__ == '__main__':
    pre = 'ABDEGCFH'
    tin = 'DBGEACHF'
    s = Solution()
    head = s.reConstructBinaryTree(pre, tin)
    post_order_print(head)  # result: DGEBHFCA
