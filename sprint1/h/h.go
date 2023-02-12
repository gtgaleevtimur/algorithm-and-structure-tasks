// Package main - решение задачи "Двоичная система".
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//str, _ := reader.ReadString('\n')
	//str = strings.TrimSpace(str)
	var n string
	fmt.Scan(&n)
	var p string
	fmt.Scan(&p)
	nSlice := strings.Split(n, "")
	pSlice := strings.Split(p, "")
	res := make([]int, 0)
	max := nSlice
	min := pSlice
	flag := int64(0)
	if len(max) < len(min) {
		max, min = min, max
	}
	max = reverse(max)
	min = reverse(min)
	for i := 0; i < len(max); i++ {
		var d1, d2 int64
		d1, _ = strconv.ParseInt(max[i], 10, 32)
		if i < len(min) {
			d2, _ = strconv.ParseInt(min[i], 10, 32)
		}
		temp := d1 + d2 + flag
		switch temp {
		case 0:
			res = append(res, 0)
		case 1:
			res = append(res, 1)
			flag = 0
		case 2:
			res = append(res, 0)
			flag = 1
		case 3:
			res = append(res, 1)
			flag = 1
		}
	}
	if flag == 1 {
		res = append(res, 1)
	}
	result := make([]string, len(res))
	for i, v := range res {
		rr := strconv.Itoa(v)
		result[i] = rr
	}
	result = reverse(result)
	out := strings.Join(result, "")
	fmt.Println(out)
}

func reverse(s []string) []string {
	if len(s) == 0 {
		return s
	}
	return append(reverse(s[1:]), s[0])
}
