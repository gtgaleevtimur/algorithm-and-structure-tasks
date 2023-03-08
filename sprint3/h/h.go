// Package main - решение задачи "Большое число".
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

	result := s(strSlice)
	fmt.Println(result)

}

func s(m []string) string {

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m)-i-1; j++ {
			first := m[j] + m[j+1]
			second := m[j+1] + m[j]
			if first < second {
				m[j], m[j+1] = m[j+1], m[j]
			}
		}
	}

	return strings.Join(m, "")
}
