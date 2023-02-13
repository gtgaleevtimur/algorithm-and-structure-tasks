// Package main - решение итоговой задачи "Ближайший ноль".
// ID 82203074
// Вычислительная сложность O(n)
// Пространственная сложность O(n)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const zero string = "0"

func main() {
	var nY int
	fmt.Scan(&nY)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strSlice := strings.Split(str, " ")
	result := make([]int, nY)
	writer := bufio.NewWriter(os.Stdout)
	digit := nY - 1
	for i := nY - 1; i >= 0; i-- {
		if strSlice[i] == zero {
			digit = 0
			result[i] = digit
			continue
		}
		digit++
		result[i] = digit
	}
	digit = nY - 1
	for i := 0; i < nY; i++ {
		temp := result[i]
		if strSlice[i] == zero {
			digit = 0
			temp = digit
			writer.WriteString(fmt.Sprintf("%d ", temp))
			continue
		}
		digit++
		if temp > digit {
			temp = digit
		}
		writer.WriteString(fmt.Sprintf("%d ", temp))
	}
	writer.Flush()
}
