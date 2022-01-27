// 面试题 16.25. LRU 缓存
// 设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。

// 它应该支持以下操作： 获取数据 get 和 写入数据 put 。

// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
// 写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

// 示例:

// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );

// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // 返回  1
// cache.put(3, 3);    // 该操作会使得密钥 2 作废
// cache.get(2);       // 返回 -1 (未找到)
// cache.put(4, 4);    // 该操作会使得密钥 1 作废
// cache.get(1);       // 返回 -1 (未找到)
// cache.get(3);       // 返回  3
// cache.get(4);       // 返回  4


type DoubleLinkedNode struct {
    Key int
    Value int
    Prev *DoubleLinkedNode
    Next *DoubleLinkedNode
}


type LRUCache struct {
    Cache map[int]*DoubleLinkedNode
    Head *DoubleLinkedNode
    Tail *DoubleLinkedNode
    Capacity int
    Size int
}


func Constructor(capacity int) LRUCache {
    res := LRUCache{
        Cache: map[int]*DoubleLinkedNode{},
        Head: &DoubleLinkedNode{},
        Tail: &DoubleLinkedNode{},
        Capacity: capacity,
        Size: 0,
    }
    // 伪头节点和伪尾节点互相指向
    res.Head.Next = res.Tail
    res.Tail.Prev = res.Head
    return res
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.Cache[key]; ok {
        // 缓存命中，移动至链表头部，表现为最近最多使用
        this.moveToHead(node)
        return node.Value
    } else {    // 未命中则 -1
        return -1
    }
}


func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.Cache[key]; ok {
        (*node).Value = value   // 缓存命中，覆盖对应缓存信息
        this.moveToHead(node)   // 变为最近最多使用
    } else {                    // 缓存未命中则创建并写入
        node := &DoubleLinkedNode{Key:key, Value:value}     // 创建
        this.Cache[key] = node      // 哈希表写入
        this.addToHead(node)        // 最近最多使用
        this.Size++                 // 缓存数量变化
        if this.Size > this.Capacity {
            // 缓存数量达到峰值，淘汰最近最少使用缓存
            removedNode := this.removedTail()
            delete(this.Cache, (*removedNode).Key)  // 哈希表信息删除
            this.Size--     // 数量变化
        }
    }
}


func (this *LRUCache) addToHead(node *DoubleLinkedNode) {
    // 缓存加入，该缓存变为最近最多使用状态
    // 首先修改该节点前后节点的指向
    (*node).Prev = this.Head
    (*node).Next = this.Head.Next
    // 首节点和次首节点指向修改，完成插入
    this.Head.Next.Prev = node
    this.Head.Next = node
}


func (this *LRUCache) removeNode(node *DoubleLinkedNode) {
    // 删除对应节点，前后节点指向修改即可
    (*node).Prev.Next = (*node).Next
    (*node).Next.Prev = (*node).Prev
}


func (this *LRUCache) moveToHead(node *DoubleLinkedNode) {
    // 缓存命中，缓存级别提升，变为最近最多使用
    this.removeNode(node)   // 从列表中删除
    this.addToHead(node)    // 然后加入至头部
}


func (this *LRUCache) removedTail() (node *DoubleLinkedNode) {
    // 当前缓存数量达到峰值，需要进行缓存淘汰，伪尾节点的前节点为此刻缓存中最近最少使用
    node = this.Tail.Prev
    this.removeNode(node)   // 删除
    return node             // 同时需要删除哈希表信息，所以返回该节点，以便获取哈希表键名
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */