package main

import "fmt"

var count int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {

	count++
	if n == nil {
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}
	return true
}
func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	fmt.Println(tree)

	tree.Insert(52)
	tree.Insert(102)
	tree.Insert(2)
	tree.Insert(23)
	tree.Insert(59)
	tree.Insert(87)
	tree.Insert(66)
	tree.Insert(310)
	tree.Insert(99)

	fmt.Println(tree.Search(23))
	fmt.Println(count)
	fmt.Println(tree.Search(13))

}
