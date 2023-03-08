// Package main - решение задачи "Печеньки".
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
	greedy, _ := reader.ReadString('\n')
	greedy = strings.TrimSpace(greedy)
	var m int
	fmt.Scan(&m)
	cookie, _ := reader.ReadString('\n')
	cookie = strings.TrimSpace(cookie)
	greedySlice := strings.Split(greedy, " ")
	cookieSlice := strings.Split(cookie, " ")
	gsd := make([]int, len(greedySlice))
	csd := make([]int, len(cookieSlice))
	for i, v := range greedySlice {
		d, _ := strconv.Atoi(v)
		gsd[i] = d
	}
	for i, v := range cookieSlice {
		d, _ := strconv.Atoi(v)
		csd[i] = d
	}
	sort.Slice(gsd, func(i, j int) bool {
		return gsd[i] > gsd[j]
	})
	sort.Ints(csd)
	count := 0
	for _, v := range gsd {
		if len(csd) != 0 && v <= csd[len(csd)-1] {
			count++
			csd[len(csd)-1] = 0
			csd = csd[:len(csd)-1]
		}
	}
	fmt.Println(count)
}
