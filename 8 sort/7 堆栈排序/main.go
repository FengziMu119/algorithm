package main

import "fmt"

func adjustMinHeap(arr []int, pos, llen int) {
	var temp int
	child := 0
	for temp = arr[pos]; 2*pos+1 <= llen; pos = child {
		child = 2*pos + 1
		if child < llen && arr[child] > arr[child+1] {
			child++
		}
		if arr[child] < temp {
			arr[pos] = arr[child]
		} else {
			break
		}
	}
	arr[pos] = temp
}

func MinHeapSort(arr []int) {
	llen := len(arr)
	for i := llen/2 - 1; i >= 0; i-- {
		adjustMinHeap(arr, i, llen-1)
	}
	for i := llen - 1; i >= 0; i-- {
		tmp := arr[0]
		arr[0] = arr[i]
		arr[i] = tmp
		adjustMinHeap(arr, 0, i-1)
	}
}

func main() {
	arr := []int{2, 3, 4, 654, 6, 48, 65, 765}
	MinHeapSort(arr)
	fmt.Println(arr)
}
