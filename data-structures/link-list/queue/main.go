package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func main() {
	q := new(Queue)

	q.addElement(23)
	q.addElement(1)
	q.removeElement()
	q.removeElement()
	q.addElement(32)

	q.display()
	val, ok := q.peek()
	if !ok {
		fmt.Println("empty Queue")
	} else {
		fmt.Println(" value at front is ", val)
	}
	ok = q.isEmpty()
	if ok {
		fmt.Println("empty Queue")
	} else {
		fmt.Println("not empty")
	}

}

func (q *Queue) addElement(val int) {
	newNode := &Node{val: val}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}

}

func (q *Queue) removeElement() int {
	if q.front == nil {
		return 0
	}
	val := q.front.val
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}
	return val

}

func (q *Queue) display() {
	for curr := q.front; curr != nil; curr = curr.next {
		fmt.Printf("%d->", curr.val)
	}
}

func (q *Queue) peek() (int, bool) {
	if q.front == nil {
		return 0, false
	}
	return q.front.val, true
}

func (q *Queue) isEmpty() bool {
	return q.front == nil
}
