package main

import (
	. "algorithm/go_type"
	"fmt"
)

/*
	给定链表 L0->L1->L2->L3...Ln-1->Ln ，把链表重新排序为 L0->Ln->L1->Ln-1->L2->Ln-2....
要求 在原来链表的基础上进行排序，即不能申请新的结点，只能修改结点的next域，不能修改数据域
*/

// 找出链表Head的中间结点，把链表从中间短程两个子链表
func findMiddleNode(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast := head //遍历链表的每次向前走2步
	slow := head //遍历链表的每次向前走1步
	slowPre := head
	for fast != nil && fast.Next != nil {
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	slowPre.Next = nil
	return slow
}

//对不带头结点的单链表翻转
func reverse(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *LNode  //前驱结点
	var next *LNode //后继结点
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func Reorder(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	cur1 := head.Next //前半部分链表的第一个结点
	mid := findMiddleNode(head.Next)
	cur2 := reverse(mid) // 后半部分链表逆序后的第一个结点
	var tmp *LNode
	// 合并链表
	for cur1.Next != nil {
		tmp = cur1.Next
		cur1.Next = cur2
		cur1 = tmp
		tmp = cur2.Next
		cur2.Next = cur1
		cur2 = tmp
	}
	cur1.Next = cur2
}
func main() {
	fmt.Println("链表排序")
	head := &LNode{}
	CreateNode(head, 8)
	PrintNode("排序前： ", head)
	Reorder(head)
	PrintNode("排序后： ", head)
}
