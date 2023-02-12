// Package main - решение задачи "Лишняя буква".
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)
	var p string
	fmt.Scan(&p)
	nSlice := strings.Split(n, "")
	pSlice := strings.Split(p, "")
	nMap := make(map[string]int)
	for _, v := range nSlice {
		nMap[v]++
	}
	pMap := make(map[string]int)
	for _, v := range pSlice {
		val, ok := nMap[v]
		if !ok {
			fmt.Println(v)
			return
		}
		pVal, ok := pMap[v]
		if ok && pVal == val {
			fmt.Println(v)
			return
		}
		pMap[v]++
	}
}
