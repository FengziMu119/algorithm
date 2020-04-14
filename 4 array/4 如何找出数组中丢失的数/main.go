package main

import "fmt"

/*
    给定一个由n-1个整数组成的未排序的数组，其中元素都是1~n 中的不同的整数，找出素组序列中缺失的数的
线性时间算法
*/
//累计求和法
func getNum(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	suma := 0
	sumb := 0
	for k, v := range arr {
		suma += v
		sumb += k
	}
	sumb = sumb + len(arr)*2 + 1
	return sumb - suma
}

// 异或法
func getNumXor(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	a := arr[0]
	b := 1
	for i := 1; i < len(arr); i++ {
		a ^= arr[i] //1^4^3^2^7^5
	}
	for j := 2; j <= len(arr)+1; j++ {
		b ^= j //1^2^3^4^5^6^7
	}
	return a ^ b
}

func main() {
	arr := []int{1, 4, 3, 2, 7, 5}
	fmt.Print("累计求和法：  ")
	fmt.Println(getNum(arr))
	fmt.Print("异或法：  ")
	fmt.Println(getNumXor(arr))

}
