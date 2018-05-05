package main

import (
	"fmt"

	"github.com/guopuke/learngo/tree"
)

// 扩展已有类型-组合形式
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}

func main() {
	// First
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)

	fmt.Println(root)

	// Second
	// nodes := []tree.Node{
	// 	{value: 3},
	// 	{},
	// 	{5, nil, &root},
	// }

	// Third
	// 自定义工厂函数
	root.Left.Right = tree.CreateTreeNode(2)

	// fmt.Println(root)

	root.Right.Left.SetValue(4)
	// root.Right.Left.Print()

	// fmt.Println()
	root.Traverse()
	fmt.Println()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		fmt.Println(nodeCount)
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
}
