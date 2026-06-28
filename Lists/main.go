package main

import "fmt"

type Node struct {
	next  *Node
	value int
}

type List struct {
	head *Node
	len  int
}

// Insert adds a new node at the beginning of the list
func (l *List) insert(v int) {
	node := Node{
		next:  nil,
		value: v,
	}

	if l.head != nil {
		node.next = l.head
	}

	l.head = &node
	l.len++
}

// Display the linked list values
func (l *List) display() {
	temp := l.head
	fmt.Printf("[")
	for temp != nil {

		fmt.Print(temp.value, ",")
		temp = temp.next
	}
	fmt.Printf("]")
}

func (l *List) Find(v int) *Node {
	for temp := l.head; temp != nil; temp = temp.next {
		if temp.value == v {
			return temp
		}
	}
	return nil
}

func (l *List) Remove(v int) {
	if l.head == nil {
		return
	}
	if l.len == 1 {
		l.head = nil
		l.len--
		return
	}
	for temp := l.head; temp != nil; temp = temp.next {
		if temp.next.value == v {
			node := temp.next
			temp.next = node.next
			node.next = nil
			l.len--
			return
		}
	}
}

func main() {
	// Initialize list
	l1 := List{}

	// Insert values
	l1.insert(22)
	l1.insert(23)
	l1.insert(24)
	l1.insert(25)
	l1.insert(27)

	// Display values
	l1.display()
}
