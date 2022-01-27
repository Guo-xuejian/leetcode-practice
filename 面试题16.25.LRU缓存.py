# 面试题 16.25. LRU 缓存
# 设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。

# 它应该支持以下操作： 获取数据 get 和 写入数据 put 。

# 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
# 写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

# 示例:

""" LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4
 """


class DoubleLinkedNode:
    def __init__(self, key=0, value=0):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None


class LRUCache:

    def __init__(self, capacity: int):
        self.cache = {}
        # 使用伪头部和伪尾部节点，这样就不用担心节点的变换了，始终拥有两个可以操作的边界节点
        self.head = DoubleLinkedNode()
        self.tail = DoubleLinkedNode()
        self.head.next = self.tail
        self.tail.prev = self.head

        self.capacity = capacity
        self.size = 0


    # 获取缓存，存在则返回，同时将该节点放到链表头部。表示最近使用，不存在则 - 1
    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1
        # 如果 key 存在，先通过哈希表定位，再移动至头部
        node = self.cache[key]
        self.moveToHead(node)
        return node.value


    # 新的缓存信息，默认合法，必定写入，存在则覆盖，不存在则创建，同时需要判定当前缓存数量是否达到峰值，是则需要删除链表尾部的节点，因为该节点对应最近最少使用（最近最多使用在头部附近）
    def put(self, key: int, value: int) -> None:
        if key not in self.cache:
            # key 不存在，创建新的节点
            node = DoubleLinkedNode(key, value)
            self.cache[key] = node  # 写入哈希表
            self.addToHead(node)    # 最新数据默认最近最多使用，放到头部
            self.size += 1          # 缓存数量增加
            if self.size > self.capacity:   # 峰值判定
                # 超过峰值限制，删除链表尾部节点
                removed = self.removeTail()
                # 哈希表也需要删除，该缓存失效
                self.cache.pop(removed.key)
                self.size -= 1  # 缓存数量减
        else:
            # key 存在，覆盖值同时移动至链表头部即可
            node = self.cache[key]
            node.value = value
            self.moveToHead(node)
    

    # 新增的数据默认为最近最多使用，放于头部
    def addToHead(self, node):
        # 新增节点指向 head 与原头部节点
        node.prev = self.head
        node.next = self.head.next
        # head 与原头部节点指向修改，完成插入
        self.head.next.prev = node
        self.head.next = node
    

    # 从链表中删除该节点
    def removeNode(self, node):
        # 修改前后节点指向
        node.prev.next = node.next
        node.next.prev = node.prev


    # 缓存访问时置于头部
    def moveToHead(self, node):
        self.remove(node)       # 删除
        self.addToHead(node)    # 置于头部
    

    # 删除最近最少使用缓存信息
    def removeTail(self):
        # 尾节点前节点即为最近最少使用，尾节点固定不会改变的
        node = self.tail.prev
        self.removeNode(node)
        # 为什么需要返回呢-------哈希表也需要删除的，节点中存储着对应键名
        return node



# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)