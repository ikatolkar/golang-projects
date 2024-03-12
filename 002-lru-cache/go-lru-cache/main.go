package main


import (
	"fmt"
)

/*
Cache
1. Implemented using queue and hashmap
2. Hash map is used for lookup, stores value->node as they key->value
3. Queue is for storing values and retaining the order
4. Every time an object is queried, if present its moved to beginning of queue, if not present inserted to beginning of queue
5. If object is inserted its inserted at head of queue
6. Cache is of size SIZE, If more than SIZE elements are inserted in cache, element is removed from tail
8. If new object is inserted and cache is not full, insert element at head

Queue
1. Implemented using linked list
2. Is double ended
3. Has head and tail

Hashmap
1. regular hashmap, value->reference_to_node
*/

const SIZE = 5

type Node struct {
	Left *Node
	Right *Node
	Value string
}

type Queue struct {
	Head *Node
	Tail *Node
	Length int
}

func NewQueue() *Queue {
	queue := Queue{}
	queue.Head = &Node{}
	queue.Tail = &Node{}
	queue.Head.Right = queue.Tail
	queue.Tail.Left = queue.Head
	queue.Length = 0
	return &queue
}

type HashMap map[string]*Node 

type Cache struct {
	Queue *Queue
	HashMap HashMap
}

func NewCache() *Cache {
	cache := Cache{}
	cache.Queue = NewQueue()
	cache.HashMap = HashMap{}
	return &cache
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Removing Node %s...\n", n.Value)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Length -= 1
	delete(c.HashMap, n.Value)
	return n
}


func (c *Cache) Add(n *Node) {
	fmt.Printf("Adding Node %s...\n", n.Value)
	n.Right = c.Queue.Head.Right
	n.Right.Left = n
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head

	c.Queue.Length += 1
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}

}

func (c *Cache) Check(s string) {
	fmt.Printf("Checking %s...\n", s)
	var node *Node
	if val, ok := c.HashMap[s]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Value: s}
	}
	c.Add(node)
	c.HashMap[s] = node
}

func (c *Cache) Cache() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	for i:=0;i<q.Length;i++{
		fmt.Printf("%s ", node.Value)
		node = node.Right
	}
	fmt.Println("\n------")
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func main() {
	fmt.Println("Starting Cache...")
	cache := NewCache()
	for _, word := range []string{"A", "B", "C", "D", "E", "B", "A", "F", "B", "C"} {
		cache.Check(word)
		cache.Display()
	}
}
