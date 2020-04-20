package main

import (
	. "algorithm/go_type"
	"fmt"
)

/*
   无序链表去除重复项，保留原来顺序
*/

// 顺序删除
func RemoveDup(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	outerCur := node.Next
	var innerCur *LNode
	var innerPre *LNode
	for ; outerCur != nil; outerCur = outerCur.Next {
		for innerCur, innerPre = outerCur.Next, outerCur; innerCur != nil; {
			if outerCur.Data == innerCur.Data {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			} else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}

// 递归删除
func RemoveDupRecursionChild(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pointer *LNode
	cur := head
	// 对以head.Next 为首的子链表删除重复结点
	head.Next = RemoveDupRecursionChild(head.Next)
	pointer = head.Next
	// 找出以head.Next为首的子链表中与head结点相同的结点并删除
	for pointer != nil {
		if head.Data == pointer.Data {
			cur.Next = pointer.Next
			pointer = cur.Next
		} else {
			pointer = pointer.Next
			cur = cur.Next
		}
	}
	return head
}
func RemoveDupRecursion(head *LNode) {
	if head == nil {
		return
	}
	head.Next = RemoveDupRecursionChild(head.Next)
}
func main() {
	head := &LNode{}
	fmt.Println("删除重复结点")
	head = CreateNodeT()
	PrintNode("删除重复结点前：", head)
	RemoveDupRecursion(head)
	PrintNode("删除重复结点后：", head)
}
