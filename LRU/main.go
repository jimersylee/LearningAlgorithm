package main

type LinkNode struct {
	key   int
	value int
	pre   *LinkNode
	next  *LinkNode
}

type LRUCache struct {
	m          map[int]*LinkNode
	capacity   int
	head, tail *LinkNode
}

// 实例化一个LRUCache
func newLRUCache(capacity int) LRUCache {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

func (this *LRUCache) Get(key int) int {
	cache := this.m
	if v, exist := cache[key]; exist {
		this.moveToHead(v)
		return v.value
	} else {
		return -1
	}
}

// 节点移动到首位
func (this *LRUCache) moveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

// Put :实现put,当存在时,相当于进行了一次get操作,需要移动到最前面;不存在时,设置一下也相当于使用,需要移动到最前面
func (this *LRUCache) Put(key int, value int) {
	head := this.head
	tail := this.tail
	cache := this.m
	// if element exists
	if v, exist := cache[key]; exist {
		// update value
		v.value = value
		this.moveToHead(v)
	} else {
		//add
		v := &LinkNode{key, value, nil, nil}
		// if size reach the capacity,delete the last element
		if len(cache) == this.capacity {
			delete(cache, tail.pre.key)
			tail.pre.pre.next = tail
			tail.pre = tail.pre.pre
		}

		v.next = head.next
		v.pre = head
		head.next = v
		head.next.pre = v
		cache[key] = v
	}
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
func (this *LRUCache) AddNode(node *LinkNode) {
	head := this.head
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node

}
