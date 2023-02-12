// Package main - решение задачи "Четные и нечетные числа".
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	str = strings.TrimSpace(str)
	strSlice := strings.Split(str, " ")
	writer := bufio.NewWriter(os.Stdout)
	digits := 0
	for _, v := range strSlice {
		d, _ := strconv.Atoi(v)
		if d == 0 {
			continue
		}
		if d == 1 {
			digits++
			continue
		}
		if r := d % 2; r == 0 {
			continue
		} else {
			digits++
			continue
		}
	}
	if digits == 0 || digits == 3 {
		writer.WriteString("WIN")
	} else {
		writer.WriteString("FAIL")
	}
	writer.Flush()
}
