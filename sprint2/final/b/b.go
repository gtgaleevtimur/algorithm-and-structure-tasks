// Package main - решение задачи "Калькулятор".
// ID 82782226
// Решение согласно условиям задачи реализовал через стек на основе слайса.
// Для проверки доступности операции создается хэш-таблица с математическими функциями.
// Вычислительная сложность: 0(n)
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

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) add(d int) {
	s.data = append(s.data, d)
}

func (s *Stack) pop() int {
	temp := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return temp
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strLine := strings.Split(str, " ")
	stack := NewStack()
	opMap := mapInit()
	for _, v := range strLine {
		if digit, err := strconv.Atoi(v); err == nil {
			stack.add(digit)
			continue
		}
		second := stack.pop()
		first := stack.pop()
		var result int
		if op, ok := opMap[v]; ok {
			result = op(first, second)
		}
		stack.add(result)
	}
	fmt.Println(stack.pop())
}

type MathFunc func(int, int) int

func mapInit() map[string]MathFunc {
	myMap := make(map[string]MathFunc)
	myMap["+"] = func(i int, i2 int) int {
		return i + i2
	}
	myMap["-"] = func(i int, i2 int) int {
		return i - i2
	}
	myMap["*"] = func(i int, i2 int) int {
		return i * i2
	}
	myMap["/"] = func(i int, i2 int) int {
		if i < 0 {
			return int(math.Floor(float64(i) / float64(i2)))
		} else {
			return i / i2
		}
	}
	return myMap
}
