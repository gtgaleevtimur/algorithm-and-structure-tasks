// Package main - решение задачи "Эффективная быстрая сортировка".
// ID 83605932
// Воспользуемся реализацией быстрой сортировки in-place, что минимизирует затраты используемой памяти.
// В качестве опоры выбирается средний элемент массива, а разбиение выполняется до тех пор,
// пока все элементы слева от опоры не станут строго меньше нее, а элементы справа не станут строго больше.
// Затем функция в рекурсии вызывает себя на левом и правом отрезке от опоры.
// Вычислительная сложность: 0(nlog(n))
// Пространственная память : 0(log(n)) так как потребуется память на стек вызовов.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Human struct {
	Login string
	Count int
	Fine  int
}

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	humans := make([]Human, n)
	for i := 0; i < n; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		strSlice := strings.Split(str, " ")
		if len(strSlice) == 3 {
			d1, _ := strconv.Atoi(strSlice[1])
			d2, _ := strconv.Atoi(strSlice[2])
			node := Human{
				Login: strSlice[0],
				Count: d1,
				Fine:  d2,
			}
			humans[i] = node
		}
	}
	quick(humans, 0, len(humans)-1)
	for _, v := range humans {
		fmt.Println(v.Login)
	}
}

func quick(arr []Human, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	mid := (left + right) / 2
	pivot := arr[mid]
	for i <= j {
		for partition(arr[i], pivot) {
			i++
		}
		for reversePartition(arr[j], pivot) {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i, j = i+1, j-1

		}
	}
	quick(arr, left, j)
	quick(arr, i, right)
}

func partition(node, pivot Human) bool {
	if node.Count == pivot.Count {
		if node.Fine == pivot.Fine {
			return node.Login < pivot.Login
		}
		return node.Fine < pivot.Fine
	}
	return node.Count > pivot.Count
}

func reversePartition(node, pivot Human) bool {
	if node.Count == pivot.Count {
		if node.Fine == pivot.Fine {
			return node.Login > pivot.Login
		}
		return node.Fine > pivot.Fine
	}
	return node.Count < pivot.Count
}
