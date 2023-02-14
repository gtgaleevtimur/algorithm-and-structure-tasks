// Package main- решение задачи "Мониторинг".
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	result := make([][]string, m)
	for i := 0; i < m; i++ {
		temp := make([]string, n)
		result[i] = temp
	}
	inx := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		str := strings.Split(line, " ")
		for u, v := range str {
			result[u][inx] = v
		}
		inx++
	}
	writer := bufio.NewWriter(os.Stdout)
	for _, value := range result {
		for _, val := range value {
			writer.WriteString(val + " ")
		}
		writer.WriteString("\n")
		writer.Flush()
	}
}
