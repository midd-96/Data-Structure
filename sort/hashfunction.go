package main

import (
	"fmt"
)

const ArraySize = 7

//hashtable will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

//bucket is a linked list in each slot
type bucket struct {
	head *bucketNode
}

//bucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

//Insert will take in akey and add it tothe hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

//Search will take i akey and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

//delete will take in akey and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

//insert will take in a key
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "Already exist")
	}
}

//search will take in a key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//delete will take in a key and delete the node
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previuosNode := b.head
	for previuosNode.next.key == k {
		//delete
		previuosNode.next = previuosNode.next.next

	}
	previuosNode = previuosNode.next
	return
}

//Init will createa bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}
	hashTable.Delete("TOKEN")

	// for _, v := range list {
	// 	fmt.Println(v)
	// }
	fmt.Println("TOKEN : ", hashTable.Search("TOKEN"))

	fmt.Println("STAN : ", hashTable.Search("STAN"))

	fmt.Println("MID : ", hashTable.Search("MID"))

}
