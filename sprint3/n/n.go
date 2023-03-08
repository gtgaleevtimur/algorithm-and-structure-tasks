// Package main - решение задачи "Клумбы".
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
	data := make([][]int, n)
	for i := 0; i < n; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		strSlice := strings.Split(str, " ")
		temp := make([]int, 2)
		for j, v := range strSlice {
			d, _ := strconv.Atoi(v)
			temp[j] = d
		}
		data[i] = temp
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	result := gen(data)
	for _, v := range result {
		res := ""
		for _, k := range v {
			res += fmt.Sprintf("%d ", k)
		}
		fmt.Println(res)
	}
}

func gen(data [][]int) [][]int {
	result := make([][]int, 0)
	first := data[0][0]
	second := data[0][1]
	for i := 0; i < len(data)-1; i++ {
		if second < data[i+1][0] {
			temp := make([]int, 2)
			temp[0] = first
			temp[1] = second
			result = append(result, temp)
			first = data[i+1][0]
			second = data[i+1][1]
		}
		if second < data[i+1][1] {
			second = data[i+1][1]
		}
	}
	temp := make([]int, 2)
	temp[0] = first
	temp[1] = second
	result = append(result, temp)
	return result
}
