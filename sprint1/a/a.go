// Package main - решение задачи "Значение функции".
package main

import (
	"bufio"
	"fmt"
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
	d1, _ := strconv.Atoi(strSlice[0])
	d2, _ := strconv.Atoi(strSlice[1])
	d3, _ := strconv.Atoi(strSlice[2])
	d4, _ := strconv.Atoi(strSlice[3])
	res := d1*(d2*d2) + d3*d2 + d4
	fmt.Println(res)
}
