// Package main - решение задачи "Комбинации".
package main

import (
	"fmt"
	"strings"
)

var myMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func main() {
	var n string
	fmt.Scan(&n)
	strSlice := strings.Split(n, "")
	gen(strSlice, "")
}

func gen(n []string, path string) {
	if len(n) == 0 {
		fmt.Print(path + " ")
		return
	}
	for _, v := range myMap[n[0]] {
		gen(n[1:], path+v)
	}
}
