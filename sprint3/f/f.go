// Package main - решение задачи "Периметр треугольника".
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strS := strings.Split(str, " ")
	temp := make([]int, len(strS))
	for i, v := range strS {
		d, _ := strconv.Atoi(v)
		temp[i] = d
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] > temp[j]
	})
	for i := 0; i < len(temp)-2; i++ {
		if temp[i] < (temp[i+1] + temp[i+2]) {
			fmt.Println(temp[i] + temp[i+1] + temp[i+2])
			break
		}
	}
}
