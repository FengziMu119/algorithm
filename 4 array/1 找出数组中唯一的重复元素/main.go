package main

import (
	"fmt"
)

/*
  数字1-1000放在含有1001个元素的数组中，其中只有唯一的一个元素值重复，其他数字均只出现一次，设计一个算法，
将重复元素找出来，要求每个数组元素只能访问一次，不适用辅助存储空间
*/

//Hash法
//方法功能：在数组中找唯一重复元素
//输入参数：arr 数组对象的引用
//返回值：重复元素的值，如果无重复则返回-1
func FindDupByHash(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	data := make(map[int]bool)
	for _, v := range arr {
		if _, ok := data[v]; ok {
			return v
		} else {
			data[v] = true
		}
	}
	return -1
}

//累加法
func FindDupBySum(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	suma := 0
	sumb := 0
	for k, v := range arr {
		suma += v
		sumb += k
	}
	sumb = sumb + len(arr)
	return sumb - suma
}

// 异或法
func FindDupByXOR(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	result := 0
	for _, v := range arr {
		result ^= v
	}
	for i := 1; i < len(arr); i++ {
		result ^= i
	}
	return result
}

//数据映射法
func FindDupByMap(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	index := 0
	i := 0
	for true {
		// 素组中的元素的值只能小于len，否则会溢出
		if arr[i] > len(arr) {
			return -1
		}
		if arr[index] < 0 {
			break
		}
		// 访问过的，通过变相反数进行标记
		arr[index] *= -1
		//index 的后继为arr[index]
		index = arr[index] * -1
		if index > len(arr) {
			fmt.Println("素组中有非法数字")
			return -1
		}
	}
	return index
}

func main() {
	arr := []int{1, 3, 4, 2, 5, 3}
	fmt.Print("Hash法:  ")
	fmt.Println(FindDupByHash(arr))
	fmt.Print("累加法:  ")
	fmt.Println(FindDupBySum(arr))
	fmt.Print("异或法:  ")
	fmt.Println(FindDupByXOR(arr))
	fmt.Print("数据映射法:  ")
	fmt.Print(FindDupByMap(arr))
}
