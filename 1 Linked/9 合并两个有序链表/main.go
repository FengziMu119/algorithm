package main

import (
	. "algorithm/go_type"
	"fmt"
)

/*
	已知两个链表head1和head2 各自有序 请把它们合并成一个链表，并要求合并后有序
*/

// 合并两个升序排列的单链表
func Merge(head1 *LNode, head2 *LNode) *LNode {
	if head1 == nil || head1.Next == nil {
		return head1
	}
	if head2 == nil || head2.Next == nil {
		return head2
	}
	cur1 := head1.Next
	cur2 := head2.Next
	var head *LNode
	var cur *LNode
	if cur1.Data.(int) > cur2.Data.(int) {
		head = head2
		cur = cur2
		cur2 = cur2.Next
	} else {
		head = head1
		cur = cur1
		cur1 = cur1.Next
	}

	for cur1 != nil && cur2 != nil {
		if cur1.Data.(int) < cur2.Data.(int) {
			cur.Next = cur1
			cur = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur = cur2
			cur2 = cur2.Next
		}
	}
	if cur1 != nil {
		cur.Next = cur1
	}
	if cur2 != nil {
		cur.Next = cur2
	}
	return head
}
func main() {
	fmt.Println("合并有序链表")
	head1 := &LNode{}
	head2 := &LNode{}
	CreateNode(head1, 4)
	CreateNode(head2, 5)
	PrintNode("head1: ", head1)
	PrintNode("head2: ", head2)
	head := Merge(head1, head2)
	PrintNode("合并后的链表：", head)
}
