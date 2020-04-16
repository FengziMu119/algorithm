package main

import (
	. "algorithm/go_type"
	"fmt"
)

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

func main() {
	head := &LNode{}
	fmt.Println("就地逆序：")
	CreateNode(head, 8)
	PrintNode("逆序前：", head)
	Reverse(head)
	PrintNode("逆序后：", head)
}
