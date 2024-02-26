package main

import (
	"container/list"
	"fmt"
)

func main() {
	c := Constructor(3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)

	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	c.Put(5, 5)

	fmt.Println("PRINT QUEUE")
	head := c.queue.Front()
	for head != nil {
		fmt.Println(head.Value)

		head = head.Next()
	}
}

type LRUCache struct {
	queue *list.List
	cap   int
	store map[int]*Node
}

type Node struct {
	data   int
	keyPtr *list.Element
}

/*
1:0,2:0,3:0
*/
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		store: make(map[int]*Node),
		queue: list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if item, ok := c.store[key]; ok {
		c.queue.MoveToFront(item.keyPtr)
		return item.data
	}

	return -1

}

func (c *LRUCache) Put(key int, value int) {
	if item, ok := c.store[key]; !ok {
		if c.cap == len(c.store) {
			back := c.queue.Back()
			c.queue.Remove(back)
			delete(c.store, back.Value.(int))
		}
		c.store[key] = &Node{data: value, keyPtr: c.queue.PushFront(key)}
	} else {
		item.data = value
		c.store[key] = item
		c.queue.MoveToFront(item.keyPtr)
	}
}
