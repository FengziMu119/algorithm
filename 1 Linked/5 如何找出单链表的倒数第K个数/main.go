package main

import (
	. "algorithm/go_type"
	"fmt"
)

/**
找出单链表中的倒数第k个元素
*/

// 快慢指针法
func FindLastK(head *LNode, k int) *LNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next
	i := 0
	for i = 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	if i < k {
		return nil
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
func main() {
	fmt.Println("寻找倒数k")
	head := &LNode{}
	CreateNode(head, 8)
	PrintNode("链表： ", head)
	fmt.Println("链表倒数第3个元素为：", FindLastK(head, 3).Data.(int))
	head = &LNode{}
	CreateNode(head, 8)
	PrintNode("旋转前： ", head)
	RotateK(head, 3)
	PrintNode("旋转后： ", head)
}

// 如何将单链表向右旋转k个位置
func RotateK(head *LNode, k int) {
	if head == nil || head.Next == nil {
		return
	}
	slow := head.Next
	fast := head.Next
	i := 0
	for i = 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	if i < k {
		return
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	tmp := slow
	slow = slow.Next
	tmp.Next = nil
	fast.Next = head.Next
	head.Next = slow
}
