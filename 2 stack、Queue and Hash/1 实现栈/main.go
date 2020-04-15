package main

import (
	"errors"
	"fmt"
)

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

func main() {
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
