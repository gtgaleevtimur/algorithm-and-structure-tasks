package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func (s *Stack) push(digit int) {
	s.data = append(s.data, digit)
}

func (s *Stack) pop() {
	if len(s.data) == 0 {
		fmt.Println("error")
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) getMax() {
	if len(s.data) == 0 {
		fmt.Println("None")
		return
	}
	temp := make([]int, len(s.data))
	copy(temp, s.data)
	sort.Ints(temp)
	fmt.Println(temp[len(temp)-1])
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
			}
		}
		if len(commandStack) == 2 {
			digit, _ := strconv.Atoi(commandStack[1])
			stack.push(digit)
		}
	}
}
