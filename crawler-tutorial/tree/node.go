package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 值接收, 和普通函数一样, 只是调整方法名位置
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 使用指针作为接收者
// 指针可以改变结构内容
// 存在 nil 指针也是可以正常调用
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " + "node. Ignored")
		return
	} else {
		node.Value = value
	}
}

// 使用自定义工厂函数
// 返回局部变量的地址
func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}
