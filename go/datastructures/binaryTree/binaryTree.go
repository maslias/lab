package binarytree

import (
	"fmt"
	"strconv"
)

type node struct {
	value  int
	height int
	left   *node
	right  *node
}

type tree struct {
	head *node
}

func (t *tree) insert(n *node) {
	if t.head == nil {
		t.head = n
		return
	}
	heightList := []*node{}
	insertProcedure(n, t.head, &heightList)

	maxHeight := len(heightList)
	for i, hl := range heightList {
		hl.height = maxHeight - i
	}
}

func insertProcedure(n *node, p *node, hl *[]*node) {
	*hl = append(*hl, p)
	if n.value <= p.value {
		if p.left == nil {
			p.left = n
			return
		}
		insertProcedure(n, p.left, hl)
	} else {
		if p.right == nil {
			p.right = n
			return
		}
		insertProcedure(n, p.right, hl)
	}
	return
}

// delete all except the head
func (t *tree) deleteByValue(v int) {
	if t.head == nil {
		return
	}

	parent, child := t.findNodeForDelete(v)
	if parent == nil {
		fmt.Println("parent missing")
		return
	}

	if child != nil {
		var state int
		if child.left == nil && child.right == nil {
			state = 1
		} else if child.left == nil {
			state = 2
		} else if child.right == nil {
			state = 2
		} else {
			state = 3
		}

		switch state {

		// no childs there
		case 1:
			if parent.left == child {
				parent.left = nil
			} else if parent.right == child {
				parent.right = nil
			}
			parent.height--

			// one child is there
		case 2:
			if parent.left == child {
				if child.left != nil {
					parent.left = child.left
				} else {
					parent.left = child.right
				}
			} else if parent.right == child {
				if child.left != nil {
					parent.right = child.left
				} else {
					parent.right = child.right
				}
			}
			parent.height--

			// both childs are there
		case 3:
			if child.left.height >= child.right.height {
				p, c := findLargestOnSmallest(child, child.left)
				c.height = child.height
				if c.left != nil {
					c.height--
					parent.height--
				}
				c.right = child.right

				if parent.left == child {
					parent.left = c
				}
				if parent.right == child {
					parent.right = c
				}

				if c.left != nil {
					p.left = c.left
				}

			} else {
				p, c := findSmallestOnLargest(child, child.right)
				c.height = child.height
				if c.right != nil {
					c.height--
					parent.height--
				}
				c.left = child.left
				if parent.left == child {
					parent.left = c
				}
				if parent.right == child {
					parent.right = c
				}

				if c.right != nil {
					p.right = c.right
				}

			}

		}

	}
}

func findLargestOnSmallest(p *node, c *node) (*node, *node) {
	if c == nil {
		return nil, nil
	}

	if c.right == nil {
		return p, c
	}
	findLargestOnSmallest(c, c.right)
	return nil, nil
}

func findSmallestOnLargest(p *node, c *node) (*node, *node) {
	if c == nil {
		return nil, nil
	}

	if c.left == nil {
		return p, c
	}
	findSmallestOnLargest(c, c.left)
	return nil, nil
}

func (t *tree) findNodeForDelete(v int) (*node, *node) {
	if t.head == nil {
		return nil, nil
	}
	list := []*node{t.head}

	var n *node

	for len(list) != 0 {
		n = list[0]
		if n.value == v {
			return n, nil
		}

		if n.left != nil && n.left.value == v {
			return n, n.left
		}
		if n.right != nil && n.right.value == v {
			return n, n.right
		}

		list = list[1:]

		if n.left != nil {
			list = append(list, n.left)
		}
		if n.right != nil {
			list = append(list, n.right)
		}
	}
	return nil, nil
}

func (t *tree) printDfs() {
	if t.head == nil {
		return
	}

	buffer := &[]int{}
	printDfsProcedure(t.head, buffer)
	fmt.Println(buffer)
}

func printDfsProcedure(n *node, buffer *[]int) {
	if n == nil {
		return
	}
	printDfsProcedure(n.left, buffer)
	*buffer = append(*buffer, n.value)
	printDfsProcedure(n.right, buffer)
	return
}

func (t *tree) printBfs() {
	if t.head == nil {
		return
	}
	list := []*node{t.head}
	out := []string{}

	for len(list) != 0 {
		n := list[0]
		out = append(out, " | "+strconv.Itoa(n.value)+":"+strconv.Itoa(n.height))
		list = list[1:]

		if n.left != nil {
			list = append(list, n.left)
		}
		if n.right != nil {
			list = append(list, n.right)
		}
	}
	fmt.Printf("%s ", out)
	fmt.Println()
}

func (t *tree) compareTrees(t2 tree) {
	if t.head == nil || t2.head == nil {
		return
	}

	fmt.Printf("compare trees: %t \n", compareTreesProcedure(t.head, t2.head))
}

func compareTreesProcedure(n1 *node, n2 *node) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if n1.value != n2.value {
		return false
	}

	return compareTreesProcedure(n1.left, n2.left) && compareTreesProcedure(n1.right, n2.right)
}

func printBfsProcedure(n *node, buffer *[]int) {
	if n == nil {
		return
	}
}

func TestStructBinaryTree() {
	t := tree{}
	t.insert(&node{value: 200})
	t.insert(&node{value: 250})
	t.insert(&node{value: 150})
	t.insert(&node{value: 100})
	t.insert(&node{value: 180})
	t.insert(&node{value: 220})
	t.insert(&node{value: 280})
	t.insert(&node{value: 290})
	// t.printDfs()
	t.printBfs()

	t.deleteByValue(250)
	t.printBfs()

	// fmt.Println(n)
}
