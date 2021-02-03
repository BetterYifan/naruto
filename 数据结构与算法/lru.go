package 数据结构与算法

// 双向链表
type LinkNode struct {
	prev, next *LinkNode
	entry      entry
}

// 理解为链表中节点的具体值
type entry struct {
	key int
	val interface{}
}

type LRUCache struct {
	// map用于加速数据查找速度
	m map[int]*LinkNode
	// map的容量，表示链表最多存储多少个节点
	cap int
	// 头、尾节点，越靠近head表示使用越多
	head, tail *LinkNode
}

// Cache初始化
func NewLRUCache(cap int) LRUCache {
	head := &LinkNode{entry: entry{}}
	tail := &LinkNode{entry: entry{}}
	head.next = tail
	tail.prev = head
	return LRUCache{
		m:    make(map[int]*LinkNode, cap),
		cap:  cap,
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) AddNode(node *LinkNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next = node
	node.next.prev = node
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) Get(key int) interface{} {
	m := this.m
	if _, ok := m[key]; ok {
		return m[key]
	}
	return nil
}

func (this *LRUCache) Put(key int, val interface{}) {
	m := this.m
	if node, ok := m[key]; ok {
		// 存在这个key，将此key在链表中前移
		node.entry.val = val
		this.MoveToHead(node)
	} else {
		// 不存在，判断容量，是否需要删除尾节点
		n := &LinkNode{
			prev:  nil,
			next:  nil,
			entry: entry{key: key, val: val},
		}
		if len(m) >= this.cap {
			// 删map,移node
			delete(m, this.tail.prev.entry.key)
			this.RemoveNode(this.tail.prev)
		}
		m[key] = n
		this.AddNode(n)
	}
}
