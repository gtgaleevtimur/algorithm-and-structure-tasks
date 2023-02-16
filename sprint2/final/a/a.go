// Package main - решение задачи "Дек".
// ID 82418315
// Для решения задачи "Дек" использовался двусвязный список, так как это решение наиболее подходит условиям задачи.
// Под двусвязный список выделится значительно меньше памяти, это связано с тем, что,
// в go массив создается на основе базовых типов и NIL значения для каждого элемента
// будут равны нулевым значениям характерными для этого типа, а не NIL. Соответсвенно с этим могут возникнуть проблемы
// к тому же мы вызовем увеличение выделенной памяти под этот массив, так как под каждый элемент будь он пустым или нет
// выделится занимаемая этим типом память. В двусязном списке основанном на указателях такого соответственно не произойдет.
// Ведь сама структура списка может быть и предполагает большее значение хранимых объектов, но не всегда это реализует.
// Вычислительная сложность: 0(1)
// Пространственная сложность: 0(n)
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
	prev  *Node
}

type Deck struct {
	head    *Node
	tail    *Node
	size    int
	maxSize int
}

func NewDeck(i int) *Deck {
	return &Deck{
		head:    nil,
		tail:    nil,
		size:    0,
		maxSize: i,
	}
}

func (d *Deck) isEmpty() bool {
	return d.size == 0
}

func (d *Deck) pushBack(value string) {
	if d.size == d.maxSize {
		fmt.Println("error")
		return
	}
	node := &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}
	if d.size == 0 {
		d.head = node
		d.tail = node
		d.size++
		return
	}
	temp := d.tail
	node.prev = temp
	temp.next = node
	d.tail = node
	d.size++
}

func (d *Deck) pushFront(value string) {
	if d.size == d.maxSize {
		fmt.Println("error")
		return
	}
	node := &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}
	if d.size == 0 {
		d.head = node
		d.tail = node
		d.size++
		return
	}
	temp := d.head
	node.next = temp
	temp.prev = node
	d.head = node
	d.size++
}

func (d *Deck) popBack() {
	if d.isEmpty() {
		fmt.Println("error")
		return
	}
	if d.size == 1 {
		value := d.head.value
		d.head = nil
		d.tail = nil
		d.size--
		fmt.Println(value)
		return
	}
	value := d.tail.value
	newTail := d.tail.prev
	newTail.next = nil
	d.tail.prev = nil
	d.tail = newTail
	d.size--
	fmt.Println(value)
}

func (d *Deck) popFront() {
	if d.isEmpty() {
		fmt.Println("error")
		return
	}
	if d.size == 1 {
		value := d.head.value
		d.head = nil
		d.tail = nil
		d.size--
		fmt.Println(value)
		return
	}
	value := d.head.value
	newHead := d.head.next
	d.head.next = nil
	newHead.prev = nil
	d.head = newHead
	d.size--
	fmt.Println(value)
}

func main() {
	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	deck := NewDeck(m)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		command := strings.Split(line, " ")
		com := command[0]
		switch len(command) {
		case 1:
			if com == "pop_back" {
				deck.popBack()
			}
			if com == "pop_front" {
				deck.popFront()
			}
		case 2:
			arg := command[1]
			if com == "push_back" {
				deck.pushBack(arg)
			}
			if com == "push_front" {
				deck.pushFront(arg)
			}
		}
	}
}
