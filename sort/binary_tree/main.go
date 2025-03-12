package main

import (
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (bt *BinaryTree) insert(data int64) *BinaryTree {
	if bt.root == nil {
		bt.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		bt.root.insert(data)
	}
	return bt
}

func (bn *BinaryNode) insert(data int64) {
	if bn == nil {
		return
	} else if data <= bn.data {
		if bn.left == nil {
			bn.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			bn.left.insert(data)
		}
	} else {
		if bn.right == nil {
			bn.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			bn.right.insert(data)
		}
	}
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	print(os.Stdout, tree.root, 0, 'M')
}
