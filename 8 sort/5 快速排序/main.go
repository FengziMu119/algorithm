package main

import "fmt"

func QuickSort(arr []int, left, right int) {
	l := left
	r := right
	pivot := arr[(left+right)/2]
	for l < r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(arr, left, r)
	}
	if right > l {
		QuickSort(arr, l, right)
	}
}
func QuickSort1(arr []int, low, high int) {
	if low >= high {
		return
	}
	i := low
	j := high
	index := arr[i]
	for i < j {
		for i < j && arr[j] >= index {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] < index {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = index
	QuickSort(arr, low, i-1)
	QuickSort(arr, i+1, high)
}
func main() {
	arr := []int{2, 3, 54, 64, 54, 32, 3, 65, 754, 2}
	QuickSort1(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
