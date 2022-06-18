package main

import (
	"fmt"
)

const Size = 5

type Node struct {
	Left  *Node
	Right *Node
	Value string
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type HashMap map[string]*Node

type Cache struct {
	Queue Queue
	Hash  HashMap
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: HashMap{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	head.Left = nil
	tail.Left = head
	tail.Right = nil
	return Queue{Head: head, Tail: tail}
}

func main() {
	fmt.Println("Welcome to cache in Golang")
	fmt.Println("START CACHE")

	cache := NewCache()

	for _, word := range []string{"parrot", "avocado", "tomato", "dog", "mango", "parrot", "pickle", "chocolate"} {
		cache.Check(word)
		cache.Display()
	}
}

func (cache *Cache) Check(word string) {
	node := &Node{}

	if val, ok := cache.Hash[word]; ok {
		node = cache.Remove(val)
	} else {
		node = &Node{Value: word}
	}
	cache.Add(node)
	cache.Hash[word] = node
}

func (cache *Cache) Display() {
	cache.Queue.Display()
}

func (queue *Queue) Display() {
	node := queue.Head.Right
	fmt.Printf("%d - [", queue.Length)
	for i := 0; i < queue.Length; i++ {
		fmt.Printf("{%s}", node.Value)
		if i < queue.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func (cache *Cache) Add(node *Node) {

	fmt.Printf("Add: %s\n", node.Value)
	head := cache.Queue.Head

	node.Right = head.Right // set the right of new node
	head.Right.Left = node  // set the left of new node's right
	node.Left = head        // set left of new node
	head.Right = node       // set head to new node

	cache.Queue.Length++ // update length of queue

	if cache.Queue.Length > Size {
		cache.Remove(cache.Queue.Tail.Left) // remove the last element
	}
}

func (cache *Cache) Remove(node *Node) *Node {
	fmt.Printf("Remove: %s\n", node.Value)

	// defining left and right nodes of the passed node
	leftNode := node.Left
	rightNode := node.Right

	leftNode.Right = rightNode
	rightNode.Left = leftNode

	cache.Queue.Length -= 1 // update length of queue

	delete(cache.Hash, node.Value) // delete the value from the hash

	// node.Right = nil
	// node.Left = nil
	return node
}
