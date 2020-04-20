package main

import (
	"fmt"
)

type LNode struct {
	data  int
	right *LNode
	down  *LNode
}

func (l *LNode) Insert(headRef *LNode, data int) *LNode {
	newNode := &LNode{data: data, down: headRef}
	headRef = newNode
	return headRef
}

//用来合并两个有序的链表
func merge(a *LNode, b *LNode) *LNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	// 把两个链表中较小的结点赋予给result
	var result *LNode
	if a.data < b.data {
		result = a
		result.down = merge(a.down, b)
	} else {
		result = b
		result.down = merge(a, b.down)
	}
	return result
}

// 把链表扁平化处理
func Flaten(root *LNode) *LNode {
	if root == nil || root.right == nil {
		return root
	}
	// 递归处理root.right链表
	root.right = Flaten(root.right)
	//把root结点对应的链表与右边的链表合并
	root = merge(root, root.right)
	return root
}

func main() {
	fmt.Println("链表扁平化")
	head := CreateNode()
	head = Flaten(head)
	PrintNode("扁平化链表后：", head)
}

// 创建链表
func CreateNode() *LNode {
	node := &LNode{}
	node = node.Insert(nil, 31)
	node = node.Insert(node, 8)
	node = node.Insert(node, 6)
	node = node.Insert(node, 3)

	node.right = node.Insert(node.right, 21)
	node.right = node.Insert(node.right, 11)

	node.right.right = node.Insert(node.right.right, 50)
	node.right.right = node.Insert(node.right.right, 22)
	node.right.right = node.Insert(node.right.right, 15)

	node.right.right.right = node.Insert(node.right.right.right, 55)
	node.right.right.right = node.Insert(node.right.right.right, 40)
	node.right.right.right = node.Insert(node.right.right.right, 39)
	node.right.right.right = node.Insert(node.right.right.right, 30)

	return node
}

// 链表打印方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	tmp := node
	for tmp != nil {
		fmt.Print(tmp.data, " ")
		tmp = tmp.down
	}
	fmt.Println()
}
