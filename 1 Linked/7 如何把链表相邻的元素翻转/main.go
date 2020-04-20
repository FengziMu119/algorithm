package main

import (
	. "algorithm/go_type"
	"fmt"
)

/*
	把链表相邻元素翻转
*/

// 就地逆序
func Reverse(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	cur := head.Next
	pre := head
	var next *LNode
	for cur != nil && cur.Next != nil {
		next = cur.Next.Next
		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next
		pre = cur
		cur = next
	}
}

func main() {
	fmt.Println("相邻元素翻转")
	head := &LNode{}
	CreateNode(head, 8)
	PrintNode("顺序输出：", head)
	Reverse(head)
	PrintNode("翻转之后： ", head)
}
