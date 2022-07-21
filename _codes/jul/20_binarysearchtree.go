package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BStree struct {
	Root *Node
}

func (bstree *BStree) Insert(data int) {
	root := bstree.Root
	if root != nil { //normal case
		root.Insert(data)
	} else { //first special case
		bstree.Root = &Node{data, nil, nil}
	}

}

func (n *Node) Insert(data int) { //for recursison
	if data < n.Val {
		if n.Left == nil {
			n.Left = &Node{data, nil, nil}
		} else {
			n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{data, nil, nil}
		} else {
			n.Right.Insert(data)
		}
	}
}

func main() {
	bstree := &BStree{}
	bstree.Insert(6)
	bstree.Insert(3)
	bstree.Insert(5)
	bstree.Insert(9)
	fmt.Println(bstree)
}
