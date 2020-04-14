package main

import "fmt"

func ShellSort(arr []int) {

	for increment := len(arr) / 2; increment > 0; increment /= 2 {
		//i序号较大的数组下标，i ,j进行比较

		for i := increment; i < len(arr); i++ {
			//进行交换
			temp := arr[i]
			//按照increment，数组从j到0进行交换比较
			for j := i - increment; j >= 0; j -= increment {
				if temp < arr[j] {
					arr[j+increment] = arr[j]
					arr[j] = temp
					temp = arr[j]
				} else { //由于数组前面按照increment已经排好序，如果temp>num[j],则不必继续比较交换下去
					break
				}
			}

		}

	}
}

func main() {
	arr := []int{58, 3, 5, 7, 23, 2, 543, 7, 11}
	ShellSort(arr)
	fmt.Println(arr)
}
