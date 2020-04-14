package main

import "fmt"

//冒泡排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
func main() {
	arr := []int{2, 34, 6, 8, 9, 5, 42}
	BubbleSort(arr)
	fmt.Println(arr)
}
