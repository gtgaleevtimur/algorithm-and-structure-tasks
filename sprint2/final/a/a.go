// Package main - решение задачи "Дек".
// ID 82947268
// Для решения задачи "Дек" использовался двусвязный список, так как это решение наиболее подходит условиям задачи.
// Под двусвязный список выделится значительно меньше памяти, это связано с тем, что,
// в go массив создается на основе базовых типов и NIL значения для каждого элемента
// будут равны нулевым значениям характерными для этого типа, а не NIL. Соответсвенно с этим могут возникнуть проблемы
// к тому же мы вызовем увеличение выделенной памяти под этот массив, так как под каждый элемент будь он пустым или нет
// выделится занимаемая этим типом память. В двусязном списке основанном на указателях такого соответственно не произойдет.
// Ведь сама структура списка может быть и предполагает большее значение хранимых объектов, но не всегда это реализует.
// Вычислительная сложность: 0(n)
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

func (d *Deck) isFull() bool {
	return d.size == d.maxSize
}

func (d *Deck) pushBack(value string) error {
	if d.isFull() {
		return errors.New("full")
	}
	d.data[d.tail] = value
	d.tail = d.nextIndex(d.tail, 1)
	d.size++
	return nil
}

func (d *Deck) pushFront(value string) error {
	if d.isFull() {
		return errors.New("full")
	}
	d.data[d.head] = value
	d.head = d.nextIndex(d.head, -1)
	d.size++
	return nil
}

func (d *Deck) popBack() (value string, err error) {
	if d.isEmpty() {
		return "", errors.New("empty")
	}
	d.tail = d.nextIndex(d.tail, -1)
	value = d.data[d.tail]
	d.size--
	return value, nil
}

func (d *Deck) popFront() (value string, err error) {
	if d.isEmpty() {
		return "", errors.New("empty")
	}
	d.head = d.nextIndex(d.head, 1)
	value = d.data[d.head]
	d.size--
	return value, nil
}

func (d *Deck) nextIndex(tempIndex, vector int) int {
	result := (tempIndex + vector) % d.maxSize
	if result == -1 {
		result = d.maxSize - 1
	}
	return result
}

type pushFunc func(value string) error
type popFunc func() (result string, err error)

func main() {
	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	deck := NewDeck(m)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	pushCommands := map[string]pushFunc{
		"push_back":  deck.pushBack,
		"push_front": deck.pushFront,
	}

	popCommands := map[string]popFunc{
		"pop_back":  deck.popBack,
		"pop_front": deck.popFront,
	}
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		command := strings.Split(line, " ")
		com := command[0]
		switch len(command) {
		case 1:
			if val, err := popCommands[com](); err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(val)
			}
		case 2:
			arg := command[1]
			if err := pushCommands[com](arg); err != nil {
				fmt.Println("error")
			}
		}
	}
}
