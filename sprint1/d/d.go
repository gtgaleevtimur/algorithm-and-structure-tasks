// Package main - решение задачи "Хаотичность погоды".
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
	slice := strings.Split(str, " ")
	if len(slice) == 1 {
		fmt.Println(1)
		return
	}
	count := 0
	sliceInt := make([]int, 0)
	for _, v := range slice {
		r, _ := strconv.Atoi(v)
		sliceInt = append(sliceInt, r)
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			if sliceInt[i] > sliceInt[i+1] {
				count++
			}
			continue
		}
		if i == n-1 {
			if sliceInt[i] > sliceInt[i-1] {
				count++
			}
			continue
		}
		if sliceInt[i] > sliceInt[i-1] && sliceInt[i] > sliceInt[i+1] {
			count++
		}
	}
	fmt.Println(count)
}
