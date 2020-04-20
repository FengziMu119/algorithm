package main

import (
	. "algorithm/go_type"
	"fmt"
)

//给定单链表中的某个结点，删除该结点
func RemoveNode(head *LNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	head.Data = head.Next.Data
	temp := head.Next
	head.Next = temp.Next
	return true
}

func main() {
	fmt.Println("删除指定结点")
	head := &LNode{}
	retNode := CreateNodeT1(head, 3)
	fmt.Printf("删除结点%v前链表\n", retNode.Data)
	result := RemoveNode(retNode)
	if result {
		PrintNode("删除该结点后链表：", head)
	}
}

func CreateNodeT1(node *LNode, get int) (retNode *LNode) {
	cur := node
	for i := 1; i < 8; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
		if i == get {
			retNode = cur
		}
	}
	return
}
