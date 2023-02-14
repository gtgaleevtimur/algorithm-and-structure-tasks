// Package main - решение задачи "Нелюбимое дело".
package main

import "fmt"

// Comment it before submitting
type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, idx int) *ListNode {
	var node *ListNode
	if idx == 0 {
		node = head.next
		return node
	}
	prevNode := NodeByIndex(head, idx-1)
	nextNode := NodeByIndex(head, idx+1)
	prevNode.next = nextNode
	return head
}
func NodeByIndex(head *ListNode, idx int) *ListNode {
	i := 0
	node := head
	for i != idx {
		node = node.next
		i++
		if i == idx {
			return node
		}
	}
	return node
}

func main() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	newHead := Solution(&node0, 1)
	fmt.Println(newHead.data, newHead.next.data)
	// result is : node0 -> node2 -> node3
}
