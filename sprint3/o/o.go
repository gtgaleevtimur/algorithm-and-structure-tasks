// Package main - решение задачи "Разность треш-индексов".
package main

import (
	"bufio"
	"fmt"
	"math"
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
	var m int
	fmt.Scan(&m)
	temp := make([]int, len(strS))
	for i, v := range strS {
		d, _ := strconv.Atoi(v)
		temp[i] = d
	}
	result := make([]int, 0)
	for i := 0; i < len(temp); i++ {
		for j := 1; j <= len(temp)-1; j++ {
			if i == j || i > j {
				continue
			}
			d := int(math.Abs(float64(temp[i] - temp[j])))
			result = append(result, d)
		}
	}
	sort.Ints(result)
	fmt.Println(result[m-1])
}
