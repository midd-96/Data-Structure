package main

import "fmt"

const AlphaSize = 26

//node
type Node struct {
	children [26]*Node
	isEnd    bool
}

//trie
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//Insert

func (t *Trie) Insert(w string) {
	wordlength := len(w)
	currentNode := t.root
	for i := 0; i < wordlength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

//Search

func (t *Trie) Search(w string) bool {
	wordlength := len(w)
	currentNode := t.root
	for i := 0; i < wordlength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main1() {
	myTrie := InitTrie()

	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	fmt.Println(myTrie.Search("argun"))
}
