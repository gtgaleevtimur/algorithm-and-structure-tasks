// Package main - решение задачи "Калькулятор".
// ID 82622246
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
	for _, v := range strLine {
		if digit, err := strconv.Atoi(v); err == nil {
			stack.add(digit)
			continue
		}
		second := stack.pop()
		first := stack.pop()
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
		stack.add(result)
	}
	fmt.Println(stack.pop())
}
