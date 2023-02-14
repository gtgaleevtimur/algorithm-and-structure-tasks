// Package main - решение задачи "Ограниченная очередь".
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Queue struct {
	data []string
	max  int
	head int
	tail int
	size int
}

func NewQueue(n int) *Queue {
	return &Queue{
		data: make([]string, n),
		max:  n,
		head: 0,
		tail: 0,
		size: 0,
	}
}

func (q *Queue) push(i string) {
	if q.size != q.max {
		q.data[q.tail] = i
		q.tail = (q.tail + 1) % q.max
		q.size++
		return
	}
	fmt.Println("error")
}

func (q *Queue) pop() {
	if q.size == 0 {
		fmt.Println("None")
		return
	}
	x := q.data[q.head]
	q.data[q.head] = ""
	q.size--
	q.head = (q.head + 1) % q.max
	fmt.Println(x)
}

func (q *Queue) peek() {
	if q.size == 0 {
		fmt.Println("None")
		return
	}
	fmt.Println(q.data[q.head])
}

func (q *Queue) sizeData() {
	fmt.Println(q.size)
}

func main() {
	var n int
	fmt.Scan(&n)
	var s int
	fmt.Scan(&s)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	queue := NewQueue(s)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		lineSlice := strings.Split(line, " ")
		if len(lineSlice) == 1 {
			switch lineSlice[0] {
			case "peek":
				queue.peek()
			case "size":
				queue.sizeData()
			case "pop":
				queue.pop()
			}
		}
		if len(lineSlice) == 2 {
			queue.push(lineSlice[1])
		}
	}
}
