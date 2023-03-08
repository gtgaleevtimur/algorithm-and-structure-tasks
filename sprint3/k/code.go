package main

import (
	"fmt"
	"reflect"
)

func merge_sort(arr []int, lf int, rg int) {
	if len(arr) <= 1 {
		return
	}
	mid := len(arr) / 2
	fmt.Println(mid)
	fmt.Println(arr)
	merge_sort(arr[:mid], lf, mid)
	merge_sort(arr[mid:], mid, rg)
	res := merge(arr, lf, mid, len(arr))
	for i := lf; i < len(arr); i++ {
		arr[i] = res[i-lf]
	}
}

func merge(arr []int, left int, mid int, right int) (result []int) {
	leftArr := arr[left:mid]
	rightArr := arr[mid:right]
	l, r := 0, 0
	for l < len(leftArr) && r < len(rightArr) {
		if leftArr[l] <= rightArr[r] {
			result = append(result, leftArr[l])
			l++
			continue
		}
		if leftArr[l] > rightArr[r] {
			result = append(result, rightArr[r])
			r++
			continue
		}
	}
	for l < len(leftArr) {
		result = append(result, leftArr[l])
		l++
	}
	for r < len(rightArr) {
		result = append(result, rightArr[r])
		r++
	}
	return
}

func main() {
	a := []int{1, 4, 9, 2, 10, 11}
	b := merge(a, 0, 3, 6)
	expected := []int{1, 2, 4, 9, 10, 11}
	if !reflect.DeepEqual(b, expected) {
		panic("WA. Merge")
	}

	c := []int{1, 4, 2, 10, 1, 2}
	merge_sort(c, 0, 6)
	fmt.Println(c)
	/*	expected = []int{1, 1, 2, 2, 4, 10}
		if !reflect.DeepEqual(c, expected) {
			panic("WA. MergeSort")
		}*/
}
