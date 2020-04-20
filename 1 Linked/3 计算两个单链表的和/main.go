package main

import (
	. "algorithm/go_type"
	"fmt"
)

/*
	给定2个单链表，链表的没个结点代表一个位数，计算两个数的和
*/

func Add(h1 *LNode, h2 *LNode) *LNode {
	if h1 == nil || h1.Next == nil {
		return h2
	}
	if h2 == nil || h2.Next == nil {
		return h1
	}
	c := 0                 //记录进位
	sum := 0               // 记录两个结点相加的值
	p1 := h1.Next          //遍历h1
	p2 := h2.Next          // 遍历h2
	resultHead := &LNode{} // 相加后链表头结点
	p := resultHead        //指向链表resultHead最后一个结点
	for p1 != nil && p2 != nil {
		p.Next = &LNode{} //指向新创建的储存相加和的结点
		sum = p1.Data.(int) + p2.Data.(int) + c
		p.Next.Data = sum % 10 // 两结点相加的和
		c = sum / 10           // 进位
		p = p.Next
		p1 = p1.Next
		p2 = p2.Next
	}
	// 链表h2比h1长，接下来只需要考虑h2剩余结点的值
	if p1 == nil {
		for p2 != nil {
			p.Next = &LNode{}
			sum = p2.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p2 = p2.Next
		}
	}
	// 链表h1比h2长，接下来只需要考虑h1剩余结点的值
	if p2 == nil {
		for p1 != nil {
			p.Next = &LNode{}
			sum = p1.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p1 = p1.Next
		}
	}
	if c == 1 {
		p.Next = &LNode{}
		p.Next.Data = 1
	}
	return resultHead
}

func main() {
	fmt.Println("链表相加")
	l1, l2 := CreateNodeT1()
	PrintNode("Head1： ", l1)
	PrintNode("Head2： ", l2)
	addResult := Add(l1, l2)
	PrintNode("相加后： ", addResult)
}

//创建链表
func CreateNodeT1() (l1 *LNode, l2 *LNode) {
	l1 = &LNode{}
	l2 = &LNode{}
	cur := l1
	for i := 1; i < 7; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i + 2
		cur = cur.Next
	}
	cur = l2
	for i := 9; i > 4; i-- {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
	return
}
