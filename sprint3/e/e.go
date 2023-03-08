// Package main - решение задачи "Покупка домов".
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
	var n, cache int
	fmt.Scanln(&n, &cache)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strS := strings.Split(str, " ")
	temp := make([]int, len(strS))
	for i, v := range strS {
		d, _ := strconv.Atoi(v)
		temp[i] = d
	}
	sort.Ints(temp)
	count := 0
	total := 0
	for _, v := range temp {
		if v+total <= cache {
			total += v
			count++
		}
	}
	fmt.Println(count)
}
