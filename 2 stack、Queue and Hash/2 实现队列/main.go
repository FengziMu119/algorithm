package main

import (
	"algorithm/go_type"
	"errors"
	"fmt"
)

type SliceQueue struct {
	arr   []int
	front int // 队列头
	rear  int //队列尾
}

//队列是否为空
func (s *SliceQueue) IsEmpty() bool {
	return s.front == s.rear
}

// 队列大小
func (s *SliceQueue) Size() int {
	return s.rear - s.front
}

// 返回队列首元素
func (s *SliceQueue) GetFront() int {
	if s.IsEmpty() {
		panic(errors.New("队列已经为空！"))
	}
	return s.arr[s.front]
}

//返回队列尾元素
func (s *SliceQueue) GetBack() int {
	if s.IsEmpty() {
		panic(errors.New("队列已经为空！"))
	}
	return s.arr[s.rear-1]
}

// 删除队列头元素
func (s *SliceQueue) DelQueue() {
	if s.rear > s.front {
		s.rear--
		s.arr = s.arr[1:]
	} else {
		panic(errors.New("队列已经为空！"))
	}
}

// 把新元素添加到队尾
func (s *SliceQueue) EnQueue(item int) {
	s.arr = append(s.arr, item)
	s.rear++
}

func SliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Slice构建队列结构")
	sliceQueue := &SliceQueue{arr: make([]int, 0)}
	sliceQueue.EnQueue(1)
	sliceQueue.EnQueue(2)
	fmt.Println("队列头元素为：", sliceQueue.GetFront())
	fmt.Println("队列尾元素为：", sliceQueue.GetBack())
	fmt.Println("队列大小为：", sliceQueue.Size())
}

type LinkedQueue struct {
	head *go_type.LNode
	end  *go_type.LNode
}

// 获取队列中元素的个数
func (l *LinkedQueue) Size() int {
	size := 0
	node := l.head
	for node != nil {
		node = node.Next
		size++
	}
	return size
}

// 入队列
func (l *LinkedQueue) EnQueue(e int) {
	node := &go_type.LNode{Data: e}
	if l.head == nil {
		l.head = node
		l.end = node
	} else {
		l.end.Next = node
		l.end = node
	}
}

// 出队列
func (l *LinkedQueue) DelQueue() {
	if l.head == nil {
		panic(errors.New("队列已经为空！"))
	}
	l.head = l.head.Next
	if l.head == nil {
		l.end = nil
	}
}

// 取得队列首元素
func (l *LinkedQueue) GetFront() int {
	if l.head == nil {
		panic(errors.New("队列已经为空！"))
	}
	return l.head.Data.(int)
}

//取得队列尾元素
func (l *LinkedQueue) GetBack() int {
	if l.end == nil {
		panic(errors.New("队列已经为空！"))
	}
	return l.end.Data.(int)
}
func LinkedMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Linked构建队列结构")
	linkedQueue := &LinkedQueue{}
	linkedQueue.EnQueue(1)
	linkedQueue.EnQueue(2)
	fmt.Println("队列首元素为：", linkedQueue.GetFront())
	fmt.Println("队列尾元素为：", linkedQueue.GetBack())
	fmt.Println("队列大小为：", linkedQueue.Size())
}
func main() {
	SliceMode()
	LinkedMode()
}
