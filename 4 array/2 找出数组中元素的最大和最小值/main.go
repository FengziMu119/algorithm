package main

import (
	"fmt"
)

/*
   给定数组a1,a2,a3....an,要求找出素组中的最大值和最小值，假设数组中的值两两不相同
*/

// 蛮力法

// 分治法
func GetMaxAndMin(arr []int) (max, min int) {
	if arr == nil || len(arr) == 0 {
		return 0, 0
	}
	max = arr[0]
	min = arr[0]
	// 两两分组，把较小的数放在左半部分，较大的放在右半部分
	for i := 0; i < len(arr)-1; i += 2 {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	// 在各个分组的左半部分找最小值
	for i := 0; i < len(arr); i += 2 {
		if arr[i] < min {
			min = arr[i]
		}
	}
	// 在各个分组的右半部分找最大值
	for i := 3; i < len(arr); i += 2 {
		if arr[i] > max {
			max = arr[i]
		}
	}
	//如果数组中的元素个数是奇数，最后一个元素被分为一组，需要特殊处理
	if len(arr)%2 == 1 {
		if max < arr[len(arr)-1] {
			max = arr[len(arr)-1]
		}
		if min > arr[len(arr)-1] {
			min = arr[len(arr)-1]
		}
	}
	return
}

// 变形的分治法
func GetMaxAndMinRe(arr []int, l, r int) (max, min int) {
	if arr == nil || len(arr) == 0 {
		return 0, 0
	}
	//求中点
	m := (l + r) / 2
	// l 和 r 之间只有一个元素
	if l == r {
		max = arr[l]
		min = arr[l]
		return
	}
	// l 和 r 之间有两个元素
	if l+1 == r {
		if arr[l] > arr[r] {
			max = arr[l]
			min = arr[r]
		} else {
			max = arr[r]
			min = arr[l]
		}
		return
	}
	// 递归计算左半部分
	lmax, lmin := GetMaxAndMinRe(arr, l, m)
	// 递归计算右半部分
	rmax, rmin := GetMaxAndMinRe(arr, m+1, r)

	// 总的最大值
	if lmax > rmax {
		max = lmax
	} else {
		max = rmax
	}
	if lmin < rmin {
		min = lmin
	} else {
		min = rmin
	}
	return
}

func main() {
	arr := []int{7, 3, 18, 24, 53, 41, 1}
	max, min := GetMaxAndMin(arr)
	fmt.Print("分治法： ")
	fmt.Print("max: ", max, " ")
	fmt.Println("min: ", min)
	max, min = GetMaxAndMinRe(arr, 0, len(arr)-1)
	fmt.Print("扩展分治法： ")
	fmt.Print("max: ", max, " ")
	fmt.Print("min: ", min)
}
