// Package main - решение задачи
// ID 83598729
// Из-за того что массив частично отсортирован, а именно две половины массива нормально упорядочены,
// и имеют точку перегиба, нарушающую порядок. Алгоритм вычисляет как отсортированы части массива рекурсивно,
// а если находит ошибку, изменяет поведение поиска.
// Вычислительная ложность :0(log(n))
// Пространственная сложность : 0(log(n))

package main

func brokenSearch(arr []int, k int) int {
	return search(arr, 0, len(arr)-1, k)
}

func search(arr []int, left, right, k int) int {
	mid := (left + right) / 2
	if arr[mid] == k {
		return mid
	}
	if right < left {
		return -1
	}
	if arr[left] <= arr[mid] {
		if arr[left] <= k && k < arr[mid] {
			return search(arr, left, mid-1, k)
		} else {
			return search(arr, mid+1, right, k)
		}
	}
	if arr[left] > arr[mid] {
		if arr[mid] < k && k <= arr[right] {
			return search(arr, mid+1, right, k)
		} else {
			return search(arr, left, mid-1, k)
		}
	}
	return -1
}

func test() {
	arr := []int{19, 21, 100, 101, 1, 4, 5, 7, 12}
	if brokenSearch(arr, 5) != 6 {
		panic("WA")
	}
}
