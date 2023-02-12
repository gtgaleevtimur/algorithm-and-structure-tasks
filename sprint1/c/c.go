// Package main - решение задачи "Соседи".
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
	var nY int
	fmt.Scan(&nY)
	var nX int
	fmt.Scan(&nX)
	matrix := make([][]int, nY)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for i := 0; i < nY; i++ {
		m := make([]int, nX)
		scanner.Scan()
		line := scanner.Text()
		values := strings.Split(line, " ")
		for y, v := range values {
			d, _ := strconv.Atoi(v)
			m[y] = d
		}
		matrix[i] = m
	}
	var y int
	fmt.Scan(&y)
	var x int
	fmt.Scan(&x)
	res := make([]int, 0)
	if y+1 < nY {
		res = append(res, matrix[y+1][x])
	}
	if y-1 >= 0 {
		res = append(res, matrix[y-1][x])
	}
	if x+1 < nX {
		res = append(res, matrix[y][x+1])
	}
	if x-1 >= 0 {
		res = append(res, matrix[y][x-1])
	}
	sort.Ints(res)
	writer := bufio.NewWriter(os.Stdout)
	str := ""
	for _, v := range res {
		r := strconv.Itoa(v)
		str += r + " "
	}
	str = strings.TrimSpace(str)
	writer.WriteString(str)
	writer.Flush()
}
