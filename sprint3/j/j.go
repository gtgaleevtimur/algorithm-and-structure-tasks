// Package main - решение задачи "Пузырек".
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
	strSlice := strings.Split(str, " ")
	temp := make([]int, n)
	for i, v := range strSlice {
		d, _ := strconv.Atoi(v)
		temp[i] = d
	}
	flag := true
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if temp[j] > temp[j+1] {
				temp[j], temp[j+1] = temp[j+1], temp[j]
				flag = true
			}
		}
		if flag {
			res := ""
			for _, v := range temp {
				res += fmt.Sprintf("%d ", v)
			}
			fmt.Println(res)
			flag = false
		}
	}
}
