package main

import (
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
func (s *SliceQueue) delQueue() {
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

func main() {
	SliceMode()
}
