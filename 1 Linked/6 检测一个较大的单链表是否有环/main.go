package main

import (
	. "algorithm/go_type"
	"fmt"
)

//快慢指针法
// 判断单链表是否有环
func IsLoop(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

// 找出环的入口点
func FindLoopNode(head *LNode, meetNode *LNode) *LNode {
	first := head.Next
	second := meetNode
	for first != second {
		first = first.Next
		second = second.Next
	}
	return first
}

func main() {
	fmt.Println("查找环")
	head := &LNode{}
	CreateNode(head, 8)
	meetNode := IsLoop(head)
	if meetNode != nil {
		fmt.Println("有环")
		loopNode := FindLoopNode(head, meetNode)
		fmt.Println("环的入口点为：", loopNode.Data.(int))
	} else {
		fmt.Println("无环")
	}
}
