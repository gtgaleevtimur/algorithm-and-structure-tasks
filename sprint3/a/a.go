// Package main - решение задачи "Генератор скобок".
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	gen(0, n, n, "")
}

func gen(temp, i, k int, prefix string) {
	if i == 0 && k == 0 {
		fmt.Println(prefix)
	}
	if i > 0 {
		gen(temp+1, i-1, k, prefix+"(")
	}
	if k > 0 && temp != 0 {
		gen(temp-1, i, k-1, prefix+")")
	}
}
