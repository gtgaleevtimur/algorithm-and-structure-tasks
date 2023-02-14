package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{
		data: make([]string, 0),
	}
}

func (s *Stack) push(st string) {
	s.data = append(s.data, st)
}

func (s *Stack) pop() {
	if len(s.data) != 0 {
		s.data = s.data[:len(s.data)-1]
	}
}

func (s *Stack) peek() string {
	if len(s.data) != 0 {
		return s.data[len(s.data)-1]
	}
	return ""
}
func main() {
	var data string
	fmt.Scan(&data)
	isCorrect(data)
}

func isCorrect(data string) {
	dMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	if data == "" {
		fmt.Println("True")
		return
	}
	str := strings.Split(data, "")
	if len(str)%2 != 0 {
		fmt.Println("False")
		return
	}
	stack := NewStack()
	for _, v := range str {
		_, ok := dMap[v]
		if ok {
			stack.push(v)
			continue
		}
		val := dMap[stack.peek()]
		if len(stack.data) != 0 && v == val {
			stack.pop()
			continue
		}
		stack.push(v)
		fmt.Println("False")
		break
	}
	if len(stack.data) == 0 {
		fmt.Println("True")
	}
}
