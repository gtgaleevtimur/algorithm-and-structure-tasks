// Package main - решение задачи "Всё наоборот".
package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

func Solution(head *ListNode) *ListNode {
	n := head
	m := n.next
	n.next = nil
	n.prev = m
	for m != nil {
		m.prev = m.next
		m.next = n
		n = m
		m = m.prev
	}
	return n
}

func main() {
	node3 := ListNode{"node3", nil, nil}
	node2 := ListNode{"node2", &node3, nil}
	node1 := ListNode{"node1", &node2, nil}
	node0 := ListNode{"node0", &node1, nil}
	node3.prev = &node2
	node2.prev = &node1
	node1.prev = &node0
	newHead := Solution(&node0)
	fmt.Println(newHead)
	// result is : newHead == node3
	// node3.next == node2
	// node2.next == node1
	// node2.prev = node3
	// node1.next == node0
	// node1.prev == node2
	// node0.prev == node1
}
