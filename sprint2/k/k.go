// Package main - решение задачи "Рекурсивные числа Фибоначчи".
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := Fib(n)
	fmt.Println(res)
}

func Fib(n int) int {
	switch {
	case n <= 1: // терминальная ветка
		return 1
	default: // рекурсивная ветка
		return Fib(n-1) + Fib(n-2)
	}
}
