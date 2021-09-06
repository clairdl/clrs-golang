package main

import "fmt"

/** CLRS Chapter 12 **/

type Bst struct {
	root *Node
}

type Node struct {
	key                 int
	parent, left, right *Node
}

type BstApi interface {
	isNil() bool
	getRoot() *Node
	insert(int)
	delete(int)
	search(int) bool
	min()
}

func newBst() *Bst {
	return &Bst{nil}
}

func (t *Bst) GetRoot() *Node {
	return t.root
}

func (bst *Bst) IsNil() bool {
	return bst.root == nil
}

func (t *Bst) Insert(key int) {
	var parent *Node
	newNode := &Node{}
	newNode.key = key

	// get parent of newNode
	for cur := t.root; cur != nil; {
		parent = cur
		if newNode.key < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	newNode.parent = parent

	if parent == nil {
		t.root = newNode // the tree was empty
	} else if newNode.key < parent.key {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
}

func (t *Bst) InOrderWalk(node *Node) []int {

	// hacky workaround so i can equality test the values
	var seq []int

	if node != nil {
		t.InOrderWalk(node.left)
		fmt.Println(node.key)
		seq = append(seq, node.key)
		t.InOrderWalk(node.right)
	}

	return seq
}
