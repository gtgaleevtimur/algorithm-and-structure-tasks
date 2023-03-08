// Package main - решение задачи "Любители конференций".
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	Count int
	Value string
}

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strS := strings.Split(str, " ")
	var m int
	fmt.Scan(&m)
	myMap := make(map[string]*Node)
	res := make([]string, 0)
	for _, v := range strS {
		if _, ok := myMap[v]; !ok {
			res = append(res, v)
			node := &Node{
				Value: v,
				Count: 1,
			}
			myMap[v] = node
		}
		myMap[v].Count++
	}
	gen(m, res, myMap)
}

func gen(m int, res []string, mym map[string]*Node) {
	sort.Strings(res)
	sort.Slice(res, func(i, j int) bool {
		return mym[res[i]].Count > mym[res[j]].Count
	})
	re := strings.Join(res[:m], " ")
	fmt.Println(re)
}
