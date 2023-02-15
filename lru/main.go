package main

import (
	"fmt"
)

func main() {
	// ["LRUCache","put","put","get","put","get","put","get","get","get"]
	// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	// [null,null,null,-1,null,-1,null,-1,-1,-1]
	// [null,null,null,1,null,-1,null,-1,3,4]
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	fmt.Println(l.Get(1))
	l.Put(3, 3)
	fmt.Println(l.Get(2))

}

//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4

type LRUCache struct {
	cap        int
	cache      map[int]*Node
	head, tail *Node
}
type Node struct {
	key, value int
	prev, next *Node
}

func Constructor(capacity int) *LRUCache {
	lru := &LRUCache{
		cap:   capacity,
		cache: make(map[int]*Node),
		head:  initNode(0, 0),
		tail:  initNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.cache[key]
	if !ok {
		return -1
	}
	// 找到后，将节点移动到头，删除原来位置节点
	l.move2Head(node)
	return node.value
}

func (l *LRUCache) Put(key int, value int) {
	// 先查找是否存在，如果key存在，则更新value
	node, ok := l.cache[key]
	if ok {
		node.value = value
		// 移动到头
		l.move2Head(node)
		return
	}
	// 不存在， 则先判断是否达到最大容量
	n := initNode(key, value)
	if len(l.cache) >= l.cap {
		// 删除尾部元素
		l.delTail()
	}
	// 增加头部节点
	l.add2Head(n)
	l.cache[key] = n
}

func initNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
		prev:  &Node{},
		next:  &Node{},
	}
}
func (l *LRUCache) delTail() {
	node := l.tail.prev
	l.delNode(node)
	delete(l.cache, node.key)
}
func (l *LRUCache) move2Head(node *Node) {
	l.delNode(node)
	l.add2Head(node)

}
func (l *LRUCache) add2Head(node *Node) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
}
func (l *LRUCache) delNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
