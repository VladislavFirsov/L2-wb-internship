package main

import (
	"fmt"
	"strings"
)

type Node struct {
	prev        *Node
	next        *Node
	data        string
	found       bool
	insensitive bool // checking for register sensitive pattern
	index       int
}

type linkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *linkedList) search(node *Node, val []string, pattern string) {
	for _, value := range val {
		if value == pattern {
			node.found = true
		}
		if strings.ToLower(string(value)) == strings.ToLower(pattern) {
			node.insensitive = true
		}
	}
}

func (l *linkedList) insert(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}

	l.length++
	node.index = l.length
}

func (l *linkedList) createNode(data string) *Node {
	return &Node{
		data: data,
	}
}

func (l *linkedList) showMatches() {
	ptr := l.head
	for ptr.next != nil {
		switch {
		case vFlag:
			l.showInversions(ptr)
		case nFlag:
			l.showStringNumber(ptr)
		case iFlag:
			l.ignoreRegister(ptr)
		default:
			if ptr.found {
				fmt.Println(ptr.data)
			}
		}
		ptr = ptr.next
	}
}

func (l *linkedList) showInversions(ptr *Node) {
	if !ptr.found {
		fmt.Println(ptr.data)
	}
}

func (l *linkedList) showStringNumber(ptr *Node) {
	if ptr.found {
		fmt.Printf("%d %s\n", ptr.index, ptr.data)
	}
}
func (l *linkedList) ignoreRegister(ptr *Node) {
	if ptr.found || ptr.insensitive {
		fmt.Println(ptr.data)
	}
}

func (l *linkedList) countMatches() int {
	counter := 0
	ptr := l.head
	for ptr.next != nil {
		if ptr.found {
			counter++
		}
		ptr = ptr.next
	}
	return counter
}

func (l *linkedList) handleAFlag(num int) error {
	ptr := l.head
	for ptr.next != nil {
		if ptr.found {
			fmt.Println(ptr.data)
			temp := ptr
			for i := 0; i < num; i++ {
				if temp.next == nil {
					return ErrOutOfRange
				}
				fmt.Println(temp.next.data)
				temp = temp.next
			}
		}
		ptr = ptr.next
	}
	return nil
}

func (l *linkedList) handleBFlag(num int) error {
	ptr := l.head
	for ptr.next != nil {
		if ptr.found {
			fmt.Println(ptr.data)
			temp := ptr
			for i := 0; i < num; i++ {
				if temp.next == nil {
					return ErrOutOfRange
				}
				fmt.Println(temp.prev.data)
				temp = temp.prev
			}
		}
		ptr = ptr.next
	}
	return nil
}

func (l *linkedList) handleCFlag(num int) error {
	ptr := l.head
	for ptr.next != nil {
		if ptr.found {
			fmt.Println(ptr.data)
			back := ptr
			forth := ptr
			for i := 0; i < num; i++ {
				if back.prev == nil || forth.next == nil {
					return ErrOutOfRange
				}
				fmt.Println(back.prev.data)
				fmt.Println(forth.next.data)
				back = back.prev
				forth = forth.next
			}
		}
		ptr = ptr.next
	}
	return nil
}
