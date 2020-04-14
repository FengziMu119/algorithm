package main

import "fmt"

// 插入排序
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}
}

func InsertSort1(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp, j := arr[i], i
		if arr[i-1] > tmp {
			for j >= 1 && arr[j-1] > tmp {
				arr[j] = arr[j-1]
				j--
			}
		}
		arr[j] = tmp
	}
}
func main() {
	arr := []int{7, 2, 3, 5, 89, 7, 3, 1}
	InsertSort(arr)
	fmt.Println(arr)
}
