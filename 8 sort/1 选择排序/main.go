package main

import "fmt"

//选择排序
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
func main() {
	arr := []int{5, 5, 6, 7, 8, 2, 4, 6, 23}
	SelectSort(arr)
	fmt.Println(arr)
}
