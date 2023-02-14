// Package main - решение задачи "Фибоначчи по модулю".
package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	a := []int{1, 1}
	d := math.Pow10(m)
	var fib int
	if n < 2 {
		fib = 1
		fmt.Println(fib)
		return
	}
	n -= 1
	for i := 0; i < n; i++ {
		s := (a[0] + a[1]) % int(d)
		a[0] = a[1]
		a[1] = s
		fib = a[1]
	}
	fmt.Println(fib)
}
