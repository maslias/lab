package tries

import (
	"bytes"
	"fmt"
	"strings"
)

type node struct {
	value  rune
	childs [26]*node
	isWord bool
}

type tries struct {
	head      *node
	wordCount int
}

func (t *tries) add(word string) {
	if t.head == nil {
		t.head = &node{}
	}

	cn := t.head
	var pos rune
	for _, r := range word {
		pos = r - 'a'
		if cn.childs[pos] == nil {
			n := &node{value: r}
			cn.childs[pos] = n
		}
		cn = cn.childs[pos]
	}

	if cn.isWord == false {
		cn.isWord = true
		t.wordCount++
	}
}

func (t *tries) findWord(word string) {
	cn := t.head
	var pos rune
	for _, r := range word {
		pos = r - 'a'
		if cn.childs[pos] == nil {
			break
		}
		cn = cn.childs[pos]
	}

	if cn == t.head {
		fmt.Println("no word there")
		return
	}

	buffer := &[]string{}
	str := bytes.Buffer{}
	getWordCmp(cn, buffer, str)

	fmt.Print(true, cn.isWord, buffer)
}

func getWordCmp(n *node, buffer *[]string, bufStr bytes.Buffer) {
	if n == nil {
		return
	}

	bufStr.WriteRune(n.value)

	for _, c := range n.childs {
		if c != nil {
			getWordCmp(c, buffer, bufStr)
		}
	}

	if n.isWord {
		w := bufStr.String()[1:]
		w = strings.TrimSpace(w)
		if w != "" {
			*buffer = append(*buffer, w)
		}
	}
}

func TestStructTries() {
	t := tries{}
	t.add("hotdog")
	t.add("hot")
	t.add("dog")
	t.add("cat")
	t.add("cats")
	t.add("catfish")
	t.add("catsfishes")
	t.add("cat")

	fmt.Printf("wordcount %d \n", t.wordCount)
	t.findWord("cat")
}
