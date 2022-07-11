package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to text search in Golang using the trie data structure")

	con := "0"

	trie := NewTrie()
	// trie.Insert("aakash")
	// fmt.Println(trie.Search("aakash"))

	var operation string // this can be add a word (a) or search a word(s)
	var word string

	for con == "0" {
		fmt.Println("Enter the operation: add(a) or search(s)")
		fmt.Scanln(&operation)

		if operation == "a" {
			// var word string

			fmt.Println("Enter the word to add")
			fmt.Scanln(&word)
			trie.Insert(word)
		} else if operation == "s" {
			// var word string

			fmt.Println("Enter the word to search")
			fmt.Scanln(&word)
			result := trie.Search(word)
			fmt.Println("Result:", result)
			if result {
				fmt.Println("Word found 😇!!")
			} else {
				fmt.Println("Word not found 🥲")
			}
		}
		fmt.Println("Do you want to continue? Press 0 to continue. Else press any key.")
		fmt.Scanln(&con)
		if con != "0" {
			break
		}
	}
}

// each node represents a letter
type Node struct {
	// char is a single letter stored. It can be a,b,c,d etc.
	char string

	// a slice of 26 nodes to store its children
	children [26]*Node
}

// a function to initialize a new node with 26 children, initially each one is nil
func NewNode(char string) *Node {
	node := &Node{char: char}
	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}
	return node
}

// trie is our actual tree that will hold all the nodes
type Trie struct {
	RootNode *Node
}

// a function to create a new trie
func NewTrie() *Trie {
	root := NewNode("R")
	return &Trie{RootNode: root}
}

// inserting a new word to the Trie
// this can be referred to as indexing
func (trie *Trie) Insert(word string) error {
	// current will keep track of our current node
	// initially it will be the root node, to start with
	current := trie.RootNode

	// pre-process the word
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	for i := 0; i < len(strippedWord); i++ {
		// to get index for a character, subtract its ascii value from that of 'a'
		index := strippedWord[i] - 'a'
		// fmt.Println("Char:", string(strippedWord[i]), " Index:", index)
		// check if node exists at index, if not, create a new one
		if current.children[index] == nil {
			current.children[index] = NewNode(string(strippedWord[i]))
		}

		// moving current to the next node
		current = current.children[index]
	}
	return nil
}

// seearching a word in the trie
func (trie *Trie) Search(word string) bool {
	// pre-process the word
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	current := trie.RootNode

	for i := 0; i < len(strippedWord); i++ {
		index := strippedWord[i] - 'a'

		if current == nil || current.children[index] == nil {
			return false
		} //else {
		// 	fmt.Println(current.children[index].char)
		// }
		current = current.children[index]
	}
	return true
}
