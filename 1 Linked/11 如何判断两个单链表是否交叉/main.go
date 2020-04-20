package main

import (
	. "algorithm/go_type"
	"fmt"
)

func IsIntersect(head1 *LNode, head2 *LNode) *LNode {
	if head1 == nil || head1.Next == nil || head2 == nil || head2.Next == nil || head1 == head2 {
		return nil
	}
	tmp1 := head1.Next
	tmp2 := head2.Next
	n1, n2 := 0, 0
	// 遍历head1，找到尾结点，同时记录head1的长度
	for tmp1.Next != nil {
		tmp1 = tmp1.Next
		n1++
	}
	//遍历head2，找到尾结点，同时记录head2的长度
	for tmp2.Next != nil {
		tmp2 = tmp2.Next
		n2++
	}
	//head1和head2是有相同的尾结点
	if tmp1 == tmp2 {
		//长链表先走|n1-n2|步
		if n1 > n2 {
			for n1-n2 > 0 {
				head1 = head1.Next
				n1--
			}
		}
		if n1 < n2 {
			for n2-n1 > 0 {
				head2 = head2.Next
				n2--
			}
		}
		for head1 != head2 {
			head1 = head1.Next
			head2 = head2.Next
		}
		return head1
	}
	return nil
}

func main() {
	fmt.Println("判断相交链表")
	head1 := &LNode{}
	retNode := CreateNodeT1(head1, 5)
	head2 := &LNode{}
	cur := head2
	for i := 1; i < 5; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
	cur.Next = retNode
	interNode := IsIntersect(head1, head2)
	if interNode == nil {
		fmt.Println("两个链表不相交")
	} else {
		fmt.Println("两个链表相交点为： ", interNode.Data)
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
