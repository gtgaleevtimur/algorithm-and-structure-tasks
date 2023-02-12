// Package main - решение задачи "Степень четырёх".
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n == 1 {
		fmt.Println("True")
		return
	}
	for n > 1 {
		if n%4 != 0 {
			fmt.Println("False")
			return
		}
		n = n / 4
	}
	fmt.Println("True")
}
