package bpm

import "errors"

type node struct {
	key   interface{}
	value interface{}
	prev  *node
	next  *node
}

type circularList struct {
	// this is a circularly linked list
	head     *node
	tail     *node
	size     int
	capacity int
}

// given a key, find the node whose value associated with it
func (c *circularList) find(key interface{}) *node {
	ptr := c.head
	for i := 0; i < c.size; i++ {
		if ptr.key == key {
			return ptr
		}
		ptr = ptr.next
	}
	return nil
}

// check if a key exists
func (c *circularList) hasKey(key interface{}) bool {
	return c.find(key) != nil
}

// insert a new key and value
func (c *circularList) insert(key interface{}, value interface{}) error {
	if c.size == c.capacity {
		return errors.New("List capacity is full")
	}

	newNode := &node{key: key, value: value}

	// no node already exists, this is the first node
	if c.size == 0 {
		newNode.next = newNode
		newNode.prev = newNode
		c.head = newNode
		c.tail = newNode
		c.size++
		return nil
	}

	// node already exixts with the same key, just update the value
	node := c.find(key)
	if node != nil {
		node.value = value
		return nil
	}

	// no node with this key exixts, but list is not empty
	newNode.next = c.head // point new node's next to head
	newNode.prev = c.tail // point new node's prev to tail
	c.tail.next = newNode // point tail's next to new node
	if c.head == c.tail {
		c.head.next = newNode // if list had one node
	}

	c.tail = newNode     // set this new node as tail
	c.head.prev = c.tail // set prev of head
	c.size++
	return nil
}

func (c *circularList) remove(key interface{}) {
	node := c.find(key)

	if node == nil {
		return
	}

	// if list had only one node, this one
	if c.size == 1 {
		c.head = nil
		c.tail = nil
		c.size--
		return
	}

	// if this node was the head of the list
	if node == c.head {
		c.head = c.head.next
	}

	// if this node was the tail of the list
	if node == c.tail {
		c.tail = c.tail.prev
	}

	// and we join the broken links here
	node.next.prev = node.prev
	node.prev.next = node.next

	c.size--
}

func (c *circularList) isFull() bool {
	return c.size == c.capacity
}
