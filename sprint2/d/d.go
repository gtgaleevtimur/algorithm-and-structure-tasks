// Package main - решение задачи "Заботливая мама".
package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, elem string) int {
	temp := head
	i := 0
	for {
		if temp == nil {
			return -1
		}
		if temp.data == elem {
			return i
		}
		i++
		temp = temp.next
	}
}

func main() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	idx := Solution(&node0, "node2")
	fmt.Println(idx)
	// result is : idx == 2
}
