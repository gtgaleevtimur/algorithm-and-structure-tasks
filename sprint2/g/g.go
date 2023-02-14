package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data []int
	max  []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
		max:  make([]int, 0),
	}
}

func (s *Stack) push(digit int) {
	dLen := len(s.data)
	mLen := len(s.max)
	if dLen == 0 {
		s.max = append(s.max, digit)
	}
	if mLen != 0 && digit >= s.max[mLen-1] {
		s.max = append(s.max, digit)
	}
	if mLen != 0 && digit < s.max[mLen-1] {
		s.max = append(s.max, s.max[mLen-1])
	}
	s.data = append(s.data, digit)
}

func (s *Stack) pop() {
	if len(s.data) == 0 || len(s.max) == 0 {
		fmt.Println("error")
		return
	}
	s.max = s.max[:len(s.max)-1]
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) getMax() {
	if len(s.data) == 0 || len(s.data) == 0 {
		fmt.Println("None")
		return
	}
	fmt.Println(s.max[len(s.max)-1])
}

func main() {
	var n int
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	stack := NewStack()
	for i := 0; i < n; i++ {
		scanner.Scan()
		command := scanner.Text()
		commandStack := strings.Split(command, " ")
		if len(commandStack) == 1 {
			switch commandStack[0] {
			case "get_max":
				stack.getMax()
			case "pop":
				stack.pop()
			default:
				continue
			}
		}
		if len(commandStack) == 2 {
			digit, _ := strconv.Atoi(commandStack[1])
			stack.push(digit)
		}
	}
}
