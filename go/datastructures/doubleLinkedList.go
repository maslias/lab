package main

import "fmt"

type node struct {
	value int
	next  *node
	prev  *node
}

type doubleLinkedList struct {
	head   *node
	tail   *node
	length int
}

func (dll *doubleLinkedList) add(n *node) {
	if dll.head == nil {
		dll.head, dll.tail = n, n
		dll.length++
		return
	}

	n.prev = dll.tail
	dll.tail.next = n
	dll.tail = n
	dll.length++
}

func (dll *doubleLinkedList) printStruct() {
	if dll.head == nil {
		fmt.Println("struct has now nodes")
		return
	}

	n := dll.head
	for range dll.length {

		fmt.Printf("%d ", n.value)
		n = n.next
	}
	fmt.Println()
}

func (dll *doubleLinkedList) removeHead() {
	if dll.head == nil {
		return
	}
	dll.length--
	if dll.length == 0 {
		dll.head = nil
		dll.tail = nil
		return
	}

	dll.head = dll.head.next
}

func (dll *doubleLinkedList) removeTail() {
	if dll.tail == nil {
		return
	}
	dll.length--
	if dll.length == 0 {
		dll.tail = nil
		dll.head = nil
		return
	}
	dll.tail = dll.tail.prev
}

func (dll *doubleLinkedList) removeByValue(v int) {
	if dll.head == nil {
		return
	}

	if dll.length == 1 {
		dll.head = nil
		dll.tail = nil
		dll.length = 0
		return
	}

	n := dll.head
	for n.value != v {
		if n.next == nil {
			return
		}
		n = n.next
	}
	if n == dll.head && n == dll.tail {
		dll.head = dll.head.next
		dll.tail = dll.tail.prev
		dll.length--
		return
	}

	if n == dll.head {
		dll.head = dll.head.next
		dll.length--
		return
	}
	if n == dll.tail {
		dll.tail = dll.tail.prev
		dll.length--
		return
	}

	n.prev.next, n.next.prev = n.next, n.prev
	dll.length--
}

func (dll *doubleLinkedList) findByValue(v int) (int, error) {
	if dll.head == nil {
		fmt.Println("struct has now nodes")
		return 0, fmt.Errorf("struct has now nodes")
	}

	n := dll.head
	for n.value != v {
		if n.next == nil {
			return 0, fmt.Errorf("does not found %d in struct", v)
		}
		n = n.next
	}

	return n.value, nil
}

func TestStructDoubleLinkedList() {
	dll := doubleLinkedList{}
	dll.add(&node{value: 1})
	dll.add(&node{value: 2})
	dll.add(&node{value: 3})
	dll.add(&node{value: 4})
	dll.add(&node{value: 5})
	dll.add(&node{value: 6})

	dll.printStruct()
	dll.removeHead()
	dll.removeTail()
	dll.removeByValue(3)
	dll.printStruct()

	found, err := dll.findByValue(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("found %d \n", found)
	}
}
