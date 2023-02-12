// Package main - решение задачи "Списочная форма".
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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	strSl := strings.Split(str, " ")
	strR := strings.Join(strSl, "")
	var p int
	fmt.Scan(&p)
	digit, _ := strconv.Atoi(strR)
	res := digit + p
	result := strconv.Itoa(res)
	ress := strings.Split(result, "")
	fmt.Println(strings.Join(ress, " "))
}
