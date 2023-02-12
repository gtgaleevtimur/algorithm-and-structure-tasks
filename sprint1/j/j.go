// Package main - решение задачи "Факторизация".
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	i := 2
	for i*i <= n {
		if n%i == 0 {
			fmt.Print(i, " ")
			n = n / i
		} else {
			i++
		}
	}
	if n > 1 {
		fmt.Print(n)
	}
}
