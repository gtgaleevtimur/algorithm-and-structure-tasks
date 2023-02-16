// Package main - решение задачи "Калькулятор".
// ID 82429346
// Решение согласно условиям задачи реализовал через стек на основе слайса.
// Вычислительная сложность: 0(1)
// Пространственная сложность: 0(n)
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strLine := strings.Split(str, " ")
	var stack []int
	for _, v := range strLine {
		if digit, err := strconv.Atoi(v); err == nil {
			stack = append(stack, digit)
			continue
		}
		first, second := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:(len(stack) - 2)]
		var result int
		switch v {
		case "+":
			result = first + second
		case "-":
			result = first - second
		case "*":
			result = first * second
		case "/":
			if first < 0 {
				temp := math.Floor(float64(first) / float64(second))
				result = int(temp)
			} else {
				result = first / second
			}
		}
		stack = append(stack, result)
	}
	fmt.Println(stack[len(stack)-1])
}
