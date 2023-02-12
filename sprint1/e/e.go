// Package main - решение задачи "Самое длинное слово".
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
	slice := strings.Split(str, " ")
	index := 0
	for i := 1; i < len(slice); i++ {
		if len([]rune(slice[i])) > len([]rune(slice[index])) {
			index = i
		}
	}
	fmt.Println(slice[index])
	fmt.Println(len([]rune(slice[index])))
}
