// Package main - решение задачи "Работа из дома".
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	nHex := strconv.FormatInt(int64(n), 2)
	fmt.Println(nHex)
}
