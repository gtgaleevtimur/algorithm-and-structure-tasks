// Package main - решение задачи "Дек".
// ID 82646037
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
	"errors"
	"fmt"
	"os"
	"strings"
)

type Deck struct {
	data    []string
	head    int
	tail    int
	size    int
	maxSize int
}

func NewDeck(i int) *Deck {
	return &Deck{
		data:    make([]string, i),
		head:    0,
		tail:    1,
		size:    0,
		maxSize: i,
	}
}

func (d *Deck) isEmpty() bool {
	return d.size == 0
}

func (d *Deck) pushBack(value string) error {
	if d.size == d.maxSize {
		return errors.New("full")
	}
	d.data[d.tail] = value
	d.tail = (d.tail + 1) % d.maxSize
	d.size++
	return nil
}

func (d *Deck) pushFront(value string) error {
	if d.size == d.maxSize {
		return errors.New("full")
	}
	d.data[d.head] = value
	d.head = (d.head - 1) % d.maxSize
	if d.head == -1 {
		d.head = d.maxSize - 1
	}
	d.size++
	return nil
}

func (d *Deck) popBack() (value string, err error) {
	if d.isEmpty() {
		return "", errors.New("empty")
	}
	d.tail = (d.tail - 1) % d.maxSize
	if d.tail == -1 {
		d.tail = d.maxSize - 1
	}
	value = d.data[d.tail]
	d.size--
	return value, nil
}

func (d *Deck) popFront() (value string, err error) {
	if d.isEmpty() {
		return "", errors.New("empty")
	}
	d.head = (d.head + 1) % d.maxSize
	value = d.data[d.head]
	d.size--
	return value, nil
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
				val, err := deck.popBack()
				if err != nil {
					fmt.Println("error")
					continue
				}
				fmt.Println(val)
			}
			if com == "pop_front" {
				val, err := deck.popFront()
				if err != nil {
					fmt.Println("error")
					continue
				}
				fmt.Println(val)
			}
		case 2:
			arg := command[1]
			if com == "push_back" {
				if err := deck.pushBack(arg); err != nil {
					fmt.Println("error")
				}
			}
			if com == "push_front" {
				if err := deck.pushFront(arg); err != nil {
					fmt.Println("error")
				}
			}
		}
	}
}
