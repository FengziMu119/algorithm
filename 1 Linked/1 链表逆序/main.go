package main

import (
	. "algorithm/go_type"
	"fmt"
)

// 就地逆序
func Reverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var pre *LNode    // 定义前驱结点
	var cur *LNode    // 定义当前结点
	next := node.Next // 后续结点

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

// 递归法
func RecursiveReverseChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := RecursiveReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

func RecursiveReverse(node *LNode) {
	firstNode := node.Next
	// 递归调用
	newHead := RecursiveReverseChild(firstNode)
	node.Next = newHead
}

func ReversePrint(node *LNode) {
	if node == nil {
		return
	}
	ReversePrint(node.Next)
	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}
}

// 插入法
func InsertReverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var cur *LNode  // 当前结点
	var next *LNode // 后继结点
	cur = node.Next.Next
	// 设置链表的第一个结点为尾结点
	node.Next.Next = nil
	// 把遍历到的结点插到头结点的后面
	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}

func main() {
	head := &LNode{}
	fmt.Println("就地逆序：")
	CreateNode(head, 8)
	PrintNode("就地逆序前：", head)
	Reverse(head)
	PrintNode("就地逆序后：", head)
	fmt.Println("递归逆序：")
	PrintNode("递归逆序前：", head)
	RecursiveReverse(head)
	fmt.Print("递归逆序后：")
	ReversePrint(head)
	fmt.Println()
	fmt.Println("插入逆序：")
	fmt.Print("插入逆序前：")
	ReversePrint(head)
	fmt.Println()
	InsertReverse(head)
	fmt.Print("插入逆序后：")
	ReversePrint(head)
}
