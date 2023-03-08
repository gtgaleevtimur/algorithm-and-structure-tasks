// Package main - решение задачи "Гардероб".
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strSlice := strings.Split(str, " ")
	fCount := 0
	result := make([]string, n)
	index := n - 1
	zIndex := 0
	for _, v := range strSlice {
		switch v {
		case "0":
			result[zIndex] = "0"
			zIndex++
		case "1":
			fCount++
		case "2":
			result[index] = "2"
			index--
		}
	}
	for fCount != 0 {
		result[index] = "1"
		fCount--
		index--
	}
	res := strings.Join(result, " ")
	fmt.Println(res)
}
