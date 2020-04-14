package main

import "fmt"

func Merge(arr []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1)
	R := make([]int, n2)
	for i, k := 0, p; i < n1; i, k = i+1, k+1 {
		L[i] = arr[k]
	}
	for i, k := 0, q+1; i < n2; i, k = i+1, k+1 {
		R[i] = arr[k]
	}
	var i, k, j int
	for i, k, j = 0, p, 0; i < n1 && j < n2; k++ {
		if L[i] > R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
	}
	if i < n1 {
		for j = i; j < n1; j, k = j+1, k+1 {
			arr[k] = L[j]
		}

	}
	if j < n2 {
		for i = j; i < n2; i, k = i+1, k+1 {
			arr[k] = R[i]
		}
	}

}
func MergeSort(arr []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(arr, p, q)
		MergeSort(arr, q+1, r)
		Merge(arr, p, q, r)
	}
}

func main() {
	arr := []int{2, 34, 5, 7, 89, 4, 32, 3}
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
