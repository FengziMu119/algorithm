package main

import "fmt"

// 排序法
//部分排序法
//类快速排序法
func FindSmallK(arr []int, low, high int, k int) int {
	i := low
	j := high
	splitElem := arr[i]
	for i < j {
		for i < j && arr[j] >= splitElem {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] <= splitElem {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = splitElem
	subArrayIndex := i - low
	if subArrayIndex == k-1 {
		return arr[i]
	} else if subArrayIndex > k-1 {
		return FindSmallK(arr, low, i-1, k)
	} else {
		return FindSmallK(arr, i+1, high, k-i-low)
	}

}
func main() {
	k := 3
	arr := []int{4, 5, 6, 3, 12, 6, 2}
	fmt.Println("第", k, "小的值为：", FindSmallK(arr, 0, len(arr)-1, k))
}
