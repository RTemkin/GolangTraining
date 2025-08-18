package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

// вставка в голову
func (l *LinkedList) insertHead(data int) {
	temp1 := &Node{data, nil}

	if l.head == nil {
		l.head = temp1
	} else {
		temp2 := l.head
		l.head = temp1
		temp1.next = temp2
	}

	l.length += 1
}

// вставка в хвост
func (l *LinkedList) insertTail(data int) {
	temp1 := &Node{data, nil}

	if l.head == nil {
		l.head = temp1
	} else {
		temp2 := l.head
		for temp2.next != nil {
			temp2 = temp2.next
		}
		temp2.next = temp1
	}

	l.length += 1
}

// вставка по номеру узла
func (l *LinkedList) insertNuber(n, data int) {

	if n == 0 {
		l.insertHead(data)
	} else if n == l.length-1 {
		l.insertTail(data)
	} else {
		temp1 := &Node{data, nil}
		temp2 := l.head

		for i := 0; i < n-1; i++ {
			temp2 = temp2.next
		}

		temp1.next = temp2.next
		temp2.next = temp1
	}

	l.length += 1
}

// удаление головы
func (l *LinkedList) deliteHead() {
	if l.head == nil {
		return
	} else {
		temp := l.head
		l.head = temp.next
		temp.next = nil
	}

	l.length -= 1
}

// удаление хвоста
func (l *LinkedList) deleteTail() {
	if l.head == nil {
		return
	} else if l.head.next == nil {
		l.head = nil

	} else {
		temp1 := l.head
		var temp2 *Node
		for temp1.next != nil {
			temp2 = temp1
			temp1 = temp1.next
		}
		temp2.next = nil
	}
	l.length -= 1
}

// удаление по номеру
func (l *LinkedList) deleteNumber(n int) {
	if n == 0 {
		l.deliteHead()
	} else if n == l.length-1 {
		l.deleteTail()
	} else {
		temp1 := l.head

		for i := 0; i < n-1; i++ {
			temp1 = temp1.next
		}

		temp2 := temp1.next
		temp1.next = temp2.next

		l.length -= 1

	}
}

// реверс
func (l *LinkedList) revers() {
	var curr, prev, next *Node
	curr = l.head
	prev = nil

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.head = prev
}

func createList(date int) *LinkedList {
	return &LinkedList{
		head:   &Node{date, nil},
		length: 1,
	}

}

func (l *LinkedList) printList() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.data, l.length)
		temp = temp.next
	}
}

func main() {

	var newList LinkedList

	newList.insertHead(10)
	newList.insertTail(20)
	newList.insertTail(30)
	newList.insertTail(40)
	newList.insertTail(50)

	newList.insertNuber(3, 35)
	newList.insertNuber(6, 55)

	newList.deleteTail()

	newList.printList()

}
