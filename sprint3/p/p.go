// Package main - решение задачи "Частичная сортировка".
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strS := strings.Split(str, " ")
	temp := make([]int, n)
	for i, v := range strS {
		d, _ := strconv.Atoi(v)
		temp[i] = d
	}
	count := 0
}
