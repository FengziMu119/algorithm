package main

import "fmt"

/*
   把一个有序数组最开始的若干个元素搬到数组的末尾，称为数组的旋转
   例如 arr {3,4,5,1,2} 为 数组 arr1{1,2,3,4,5} 的旋转数组
*/

// 二分查找 mid = (low +high)/2
// 1） 如果arr[mid]<arr[mid-1]，则arr[mid]一定是最小值
// 2） 如果arr[mid+1]<arr[mid]，则arr[mid+1]一定是最小值
// 3） 如果arr[high]>arr[mid]，则最小值一定在数组的左半部分
// 4） 如果arr[mid]>arr[low]，则最小值一定在数组的右半部分
// 5） 如果arr[low]==arr[mid]且arr[high]==arr[mid]，无法区分左右找最小值
func getMinPara(arr []int, low, high int) int {
	// 如果旋转个数为0，即没有旋转，单独处理，直接返回数组头元素
	if high < low {
		return arr[0]
	}
	// 只剩下一个元素一定是最小值
	if high == low {
		return arr[low]
	}
	mid := low + (high-low)>>1
	//判断是否arr[mid]为最小值
	if arr[mid] < arr[mid-1] {
		return arr[mid]
	} else if arr[mid+1] < arr[mid] { //判断是否arr[mid+1]为最小值
		return arr[mid+1]
	} else if arr[high] > arr[mid] {
		//最小值一定在数组的左半部
		return getMinPara(arr, low, mid-1)
	} else if arr[mid] > arr[low] {
		//最小值一定在数组右半部分
		return getMinPara(arr, mid+1, high)
	} else {
		left := getMinPara(arr, low, mid-1)
		right := getMinPara(arr, mid+1, high)
		if left < right {
			return left
		} else {
			return right
		}
	}
}

func getMin(arr []int) int {
	if arr == nil {
		return -1
	} else {
		return getMinPara(arr, 0, len(arr)-1)
	}
}

func swap(arr []int, low, high int) {
	//交换数组low到high的内容
	for ; low < high; low, high = low+1, high-1 {
		arr[low], arr[high] = arr[high], arr[low]
	}
}
func main() {
	fmt.Print("找出最小值:  ")
	arr := []int{5, 6, 1, 2, 3, 4}
	fmt.Println(getMin(arr))
}
