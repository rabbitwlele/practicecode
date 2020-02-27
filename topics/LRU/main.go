package main

import (
	"fmt"
)

type Node struct {
	key   int
	value int
	next  *Node
	front *Node
}

type LRUCache struct {
	capacity int
	mapping  map[int]*Node
	tail     *Node
	head     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		mapping:  make(map[int]*Node),
		tail:     nil,
		head:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.mapping[key]; ok {
		this.del(node.key)
		this.add(node.key, node.value)
		return node.value
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if node, ok := this.mapping[key]; !ok {
		if this.Len() == this.capacity {
			this.del(this.tail.key)
		}
		this.add(key, value)
	} else {
		this.del(node.key)
		this.add(key, value)
	}
}

func (this *LRUCache) Len() int {
	return len(this.mapping)
}
func (this *LRUCache) del(key int) {
	node, ok := this.mapping[key]
	if !ok {
		return
	}
	delete(this.mapping, key)
	if node.front != nil {
		node.front.next = node.next
	}
	if node.next != nil {
		node.next.front = node.front
	}
	if node == this.head {
		this.head = node.next
	}
	if node == this.tail {
		this.tail = node.front
	}
}

func (this *LRUCache) add(key, value int) {
	node := &Node{key, value, this.head, nil}
	this.mapping[key] = node
	if this.head != nil {
		this.head.front = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cache := Constructor(1)
	cache.Put(1, 1)
	fmt.Println(cache.mapping, cache.tail, cache.head)
	cache.Put(2, 2)
	fmt.Println(cache.mapping, cache.tail, cache.head)
	cache.Put(3, 3)
	fmt.Println(cache.mapping, cache.tail, cache.head)
	cache.Put(4, 4)
	fmt.Println(cache.mapping, cache.tail, cache.head)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(4))
	fmt.Println(cache.mapping, cache.tail, cache.head)
	fmt.Println(cache.Get(2))
	fmt.Println(cache.mapping, cache.tail, cache.head)
}
