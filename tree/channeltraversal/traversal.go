package main

import (
	"fmt"

	"github.com/guopuke/learngo/tree"
)

// 一个包的方法可以分别放在多个文件内
func (node *tree.Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
	// if node == nil {
	// 	return
	// }
	//
	// node.Left.Traverse()
	// node.Print()
	// node.Right.Traverse()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node != nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)

}
