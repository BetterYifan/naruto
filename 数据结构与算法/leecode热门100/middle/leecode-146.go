package middle

import "container/list"

type LRUCache struct {
	capacity int

	lru   *list.List
	cache map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		lru:      list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, hit := this.cache[key]; hit {
		this.lru.MoveToFront(v)
		return v.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.cache == nil {
		this.lru = list.New()
		this.cache = make(map[int]*list.Element)
	}
	if v, hit := this.cache[key]; hit {
		this.cache[key].Value = &entry{key: key, value: value}
		this.lru.MoveToFront(v)
		return
	}

	ele := this.lru.PushFront(&entry{key: key, value: value})
	this.cache[key] = ele

	if this.lru.Len() > this.capacity {
		ele := this.lru.Back()
		this.lru.Remove(ele)
		delete(this.cache, ele.Value.(*entry).key)
	}
}
