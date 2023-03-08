// Package main - решение задачи "Подпоследовательность".
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m string
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Println(gen(n, m))
}

func gen(n, m string) string {
	index := 0
	for _, v := range []rune(n) {
		temp := []rune(m)
		index = strings.IndexRune(string(temp[index:]), v)
		if index == -1 {
			return "False"
		}
	}
	return "True"
}
