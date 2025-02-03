package stack

import "fmt"

type node struct {
	value int
	prev  *node
}

type stack struct {
	tail   *node
	length int
}

func (s *stack) push(n *node) {
	s.length++
	if s.tail == nil {
		s.tail = n
		return
	}

	n.prev = s.tail
	s.tail = n
}

func (s *stack) pop() {
	if s.length == 0 {
		return
	}

	if s.length == 1 {
		s.tail = nil
		s.length--
		return
	}

	s.length--
	s.tail = s.tail.prev
}

func (s *stack) printStruct() {
	if s.length == 0 {
		fmt.Println("stack has no nodes")
		return
	}

	n := s.tail
	for range s.length {
		fmt.Printf(" %d ", n.value)
		n = n.prev
	}

	fmt.Println()
}

func (s *stack) findByValue(v int) (int, error) {
	if s.tail == nil {
		fmt.Println("struct has now nodes")
		return 0, fmt.Errorf("struct has now nodes")
	}

	n := s.tail
	for n.value != v {
		if n.prev == nil {
			return 0, fmt.Errorf("does not found %d in struct", v)
		}
		n = n.prev
	}

	return n.value, nil
}

func TestStructStack() {
	s := stack{}
	s.push(&node{value: 1})
	s.push(&node{value: 2})
	s.push(&node{value: 3})
	s.push(&node{value: 4})
	s.push(&node{value: 5})
	s.push(&node{value: 6})

	s.printStruct()
	s.pop()
	s.printStruct()
	s.pop()
	s.printStruct()
	s.push(&node{value: 7})
	s.printStruct()
	s.pop()
	s.printStruct()

	found, err := s.findByValue(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("found %d \n", found)
	}
}
