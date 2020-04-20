package main

import (
	. "algorithm/go_type"
	"fmt"
)

func reverse(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *LNode
	var next *LNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func ReverseK(head *LNode, k int) {
	if head == nil || head.Next == nil {
		return
	}
	pre := head
	begin := head.Next
	var end *LNode
	var pNext *LNode
	for begin != nil {
		end = begin
		for i := 1; i < k; i++ {
			if end.Next != nil {
				end = end.Next
			} else {
				return
			}
		}
		pNext = end.Next
		end.Next = nil
		pre.Next = reverse(begin)
		begin.Next = pNext
		pre = begin
		begin = pNext
	}
}

func main() {
	fmt.Println("k元素翻转")
	head := &LNode{}
	CreateNode(head, 8)
	PrintNode("顺序输出： ", head)
	ReverseK(head, 3)
	PrintNode("逆序输出： ", head)
}
