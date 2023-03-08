// Package main - решение задачи "Два велосипеда".
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strSlice := strings.Split(str, " ")
	digits := make([]int, n)
	var coast int
	fmt.Scan(&coast)
	for i, v := range strSlice {
		d, _ := strconv.Atoi(v)
		digits[i] = d
	}
	fmt.Print(binSearch(digits, coast, 0, len(digits)), " ")
	fmt.Print(binSearch(digits, coast*2, 0, len(digits)))
}

func binSearch(cache []int, coast, left, right int) int {
	if right <= left {
		return -1
	}
	middle := (left + right) / 2
	if middle != 0 {
		if cache[middle] >= coast && cache[middle-1] < coast {
			return middle + 1
		}
	} else {
		if cache[middle] >= coast {
			return middle + 1
		}
	}
	if coast <= cache[middle] {
		return binSearch(cache, coast, left, middle)
	}
	return binSearch(cache, coast, middle+1, right)
}
