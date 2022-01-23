class BinaryTreeNode:
    def __init__(self, data):
        self.val = data
        self.left = None
        self.right = None


class BST():    # BinarySearchTree
    def __init__(self, node_list):
        self.root = None
        for node in node_list:
            self.insrt(node)
    
    # 搜索
    def search(self, data):
        bt = self.root
        while bt:
            value = bt.val
            if value == data:
                return True
            elif value > data:  # 大左小右
                bt = bt.right
            else:
                bt = bt.left
        return False
    
    # 插入
    def insert(self, data):
        bt = self.root
        if not bt:  # 空树创建
            self.root = BinaryTreeNode(data)
            return
        while True:
            value = bt.val
            if data < value:
                if not bt.left:
                    bt.left = BinaryTreeNode(data)
                    return
                bt = bt.left
            elif data > value:
                if not bt.right:
                    bt.right = BinaryTreeNode(data)
                    return
                bt = bt.right
            else:
                # bt.val = data
                return
    
    # 前序遍历
    def pre_traversal(self, root):
        res = []
        if not root:
            return res
        res.append(root.val)
        self.pre_traversal(root.left)
        self.pre_traversal(root.right)
        return res
    
    # 中序遍历
    def mid_traversal(self, root):
        res = []
        if not root:
            return res
        self.pre_traversal(root.left)
        res.append(root.val)
        self.pre_traversal(root.right)
        return res
    
    # 后续遍历
    def mid_traversal(self, root):
        res = []
        if not root:
            return res
        self.pre_traversal(root.left)
        self.pre_traversal(root.right)
        res.append(root.val)
        return res

    # 删除
    def delete(self, data):
        # 删除操作需要记住父节点来控制指向
        parent, node = None, self.root
        if not node:
            print("The Tree is Null")
            return False
        
        # 先查看该元素是否存在
        while node and node.val != data:
            parent = node
            if data < node.val:
                node = node.left
            else:
                node = node.right
            if not node:
                return False
        
        if parent != None and parent.left == node:  # 不为根节点且为左子节点
            if not node.left and not node.right:    # 无子节点，直接置空
                parent.left = None
                del node
            elif node.left and not node.right:      # 存在左子节点
                parent.left = node.left
                del node
            elif not node.left and node.right:      # 存在右子节点
                parent.left = node.right
                del node
            else:                                   # 左右子节点均存在
                # 查找左子树的最右子节点，将其作为新的起始左子节点，新左子节点的左右子节点指向原左子节点的左右节点，要注意这个新左子节点很有可能存在其自身的左子节点，最后需要处理一下
                r = node.left           # 子
                p = node                # 父
                while r.right:          # 找到左子树的最右子节点
                    p = r
                    r = r.right
                l = r.left              # 别忘了该节点可能存在左子节点
                r.left = node.left      # 覆盖
                r.right = node.right
                p.right = l             # 父节点的右子节点覆盖为新节点的左子节点
                parent.left = r         # 修改指向
                del node
        elif parent != None and parent.right == node:   # 不为根节点且为右子节点，处理同上
            if not node.left and not node.right:
                parent.right = None
                del node
            elif node.left and not node.right:
                parent.right = node.left
                del node
            elif not node.left and node.right:
                parent.right = node.right
                del node
            else:
                r = node.left
                p = node
                while r.right:
                    p = r
                    r = r.right
                l = r.left
                r.left = node.left
                r.right = node.right
                p.right = l
                parent.right = r
                del node
        elif not parent:                                # node 为根节点
            r = node.left
            p = node
            while r.right:
                p = r
                r = r.right
            l = r.left
            r.left = node.left
            r.right = node.right
            p.right = l
            self.root = r   # 修改根节点指向
            del node
