// Package main - решение задачи "Списочная очередь".
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	value string
	next  *Node
}

type ListQueue struct {
	size int
	head *Node
	tail *Node
}

func NewLQ() *ListQueue {
	return &ListQueue{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (lq *ListQueue) get() {
	if lq.head == nil {
		fmt.Println("error")
		return
	}
	x := lq.head.value
	nextNode := lq.head.next
	lq.head = nextNode
	lq.size--
	fmt.Println(x)
}

func (lq *ListQueue) put(v string) {
	lq.size++
	node := &Node{
		value: v,
		next:  nil,
	}
	if lq.head == nil {
		lq.head = node
		lq.tail = node
		return
	}
	prev := lq.tail
	prev.next = node
	lq.tail = node
}

func (lq *ListQueue) sizeOf() {
	fmt.Println(lq.size)
}

func main() {
	var n int
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	lq := NewLQ()
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		command := strings.Split(line, " ")
		if len(command) == 1 {
			switch command[0] {
			case "get":
				lq.get()
			case "size":
				lq.sizeOf()
			}
		}
		if len(command) == 2 {
			lq.put(command[1])
		}
	}
}
