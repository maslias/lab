package queue

import "fmt"

type node struct {
	value int
	next  *node
}

type queue struct {
	head   *node
	tail   *node
	length int
}

func (q *queue) enqueue(n *node) {
	q.length++
	if q.head == nil {
		q.head = n
		q.tail = n
		return
	}
	q.tail.next = n
	q.tail = n
}

func (q *queue) dequeue() {
	if q.length == 1 {
		q.head = nil
		q.tail = nil
		q.length--
		return
	}

	q.head = q.head.next
	q.length--
}

func (q *queue) printStruct() {
	if q.length == 0 {
		fmt.Println("queue is emptry")
		return
	}

	n := q.head
	for range q.length {
		fmt.Printf("%d ", n.value)
		n = n.next
	}
	fmt.Println()
}

func (q *queue) findByValue(v int) (int, error) {
	if q.head == nil {
		fmt.Println("struct has now nodes")
		return 0, fmt.Errorf("struct has now nodes")
	}

	n := q.head
	for n.value != v {
		if n.next == nil {
			return 0, fmt.Errorf("does not found %d in struct", v)
		}
		n = n.next
	}

	return n.value, nil
}

func TestStructQueue() {
	q := queue{}
	q.enqueue(&node{value: 1})
	q.enqueue(&node{value: 2})
	q.enqueue(&node{value: 3})
	q.enqueue(&node{value: 4})
	q.enqueue(&node{value: 5})
	q.enqueue(&node{value: 6})
	q.printStruct()
	q.dequeue()
	q.printStruct()
	q.dequeue()
	q.printStruct()
	q.enqueue(&node{value: 7})
	q.printStruct()
	q.dequeue()
	q.printStruct()
	q.dequeue()
	q.printStruct()
	q.dequeue()
	q.printStruct()
	q.dequeue()
	q.printStruct()
}
