package main

import (
	"algorithm/go_type"
	"errors"
	"fmt"
)

// 切片方式实现
type SliceStack struct {
	arr       []int
	stackSize int
}

func (s *SliceStack) IsEmpty() bool {
	return s.stackSize == 0
}

func (s *SliceStack) Size() int {
	return s.stackSize
}

func (s *SliceStack) Top() int {
	if s.IsEmpty() {
		panic(errors.New("栈已经为空！"))
	}
	return s.arr[s.stackSize-1]
}

func (s *SliceStack) Pop() int {
	if s.stackSize > 0 {
		s.stackSize--
		ret := s.arr[s.stackSize]
		s.arr = s.arr[:s.stackSize]
		return ret
	}
	panic(errors.New("栈已经为空！"))
}

func (s *SliceStack) Push(t int) {
	s.arr = append(s.arr, t)
	s.stackSize++
}

func SliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Slice构建栈结构")
	sliceStack := &SliceStack{arr: make([]int, 0)}
	sliceStack.Push(1)
	sliceStack.Push(2)
	sliceStack.Push(3)
	sliceStack.Push(4)
	sliceStack.Push(5)
	fmt.Println("栈定元素为：", sliceStack.Top())
	fmt.Println("栈的大小为：", sliceStack.Size())
	fmt.Println("弹出元素为：", sliceStack.Pop())
	fmt.Println("弹栈成功: ", sliceStack.Size())
}

// 链表结构实现
type LinkedStack struct {
	head *go_type.LNode
}

func (l *LinkedStack) IsEmptyL() bool {
	return l.head.Next == nil
}

func (l *LinkedStack) SizeL() int {
	size := 0
	node := l.head.Next
	for node != nil {
		node = node.Next
		size++
	}
	return size
}

func (l *LinkedStack) PushL(e int) {
	node := &go_type.LNode{
		Data: e,
		Next: l.head.Next,
	}
	l.head.Next = node
}

func (l *LinkedStack) PopL() int {
	tmp := l.head.Next
	if tmp != nil {
		l.head.Next = tmp.Next
		return tmp.Data.(int)
	}
	panic(errors.New("栈已经为空！"))
}

func (l *LinkedStack) TopL() int {
	if l.head.Next != nil {
		return l.head.Next.Data.(int)
	}
	panic(errors.New("栈已经为空！"))
}

func LinkedMode() {
	fmt.Println("Linked构建栈结构")
	linkedStack := &LinkedStack{head: &go_type.LNode{}}
	linkedStack.PushL(1)
	linkedStack.PushL(2)
	linkedStack.PushL(3)
	linkedStack.PushL(4)
	linkedStack.PushL(5)
	fmt.Println("栈定元素为：", linkedStack.TopL())
	fmt.Println("栈的大小为：", linkedStack.SizeL())
	fmt.Println("弹出元素为：", linkedStack.PopL())
	fmt.Println("弹栈成功: ", linkedStack.SizeL())
}

func main() {
	SliceMode()
	LinkedMode()
}
