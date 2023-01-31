package main

import (
	. "algorithm/go_type"
	"fmt"
	"math"
)

// 蛮力法
func minDistance(arr []int, num1, num2 int) int {
	if arr == nil || len(arr) < 1 {
		return math.MaxInt64
	}
	minDis := math.MaxInt64
	dist := 0
	for i, v := range arr {
		if v == num1 {
			for j, v := range arr {
				if v == num2 {
					dist = Abs(i - j)
				}
			}
			if dist < minDis {
				minDis = dist
			}
		}
	}
	return minDis
}

// 动态规划法
func minDynDistance(arr []int, num1, num2 int) int {
	if arr == nil || len(arr) < 1 {
		return math.MaxInt64
	}
	lastPost1 := 0 //上次遍历到num1的位置
	lastPost2 := 0 // 上次遍历到num1的位置
	minDis := math.MaxInt64
	for i, v := range arr {
		if v == num1 {
			lastPost1 = i
			if lastPost2 > 0 {
				minDis = Min(minDis, lastPost1-lastPost2)
			}
		}
		if v == num2 {
			lastPost2 = i
			if lastPost1 > 0 {
				minDis = Min(minDis, lastPost2-lastPost1)
			}
		}
	}
	return minDis
}

func main() {
	arr := []int{4, 5, 6, 7, 8, 9, 6, 4, 2, 8, 8, 6, 4, 4}
	num1 := 4
	num2 := 8
	fmt.Println("蛮力法")
	fmt.Println(minDistance(arr, num1, num2))
	fmt.Println("动态规划法")
	fmt.Println(minDynDistance(arr, num1, num2))
}
