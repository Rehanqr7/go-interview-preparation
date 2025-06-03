package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkList struct {
	head *Node
}

func main() {
	ll := new(LinkList)

	ll.addElement(2)
	ll.addElement(4)
	ll.addElement(45)
	ll.addElement(3)
	ll.addElement(23)

	ll.display()

}

func (h *LinkList) addElement(val int) {

	newNode := &Node{val: val}

	if h.head == nil {
		h.head = newNode
		return
	}
	current := h.head

	for current.next != nil {
		current = current.next
	}
	current.next = newNode

}

func (h *LinkList) display() {

	current := h.head
	for current != nil {
		fmt.Printf("%d->", current.val)
		current = current.next
	}
}
