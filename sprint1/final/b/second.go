// Package main - решение итоговой задачи "Ловкость рук".
// ID 82171923
// Вычислительная сложность O(1)
// Пространственная сложность O(1)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var k int
	fmt.Scan(&k)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	digits := make([]int, 9)
	for i := 0; i < 4; i++ {
		scanner.Scan()
		line := scanner.Text()
		values := strings.Split(line, "")
		for _, v := range values {
			if d, err := strconv.Atoi(v); err == nil {
				digits[d-1] += d
			}
		}
	}
	result := 0
	for i, v := range digits {
		temp := v / (i + 1)
		if 0 < temp && temp <= k*2 {
			result++
		}
	}
	fmt.Println(result)
}
