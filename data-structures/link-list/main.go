package main

import "fmt"

type Node struct {
	val  int
	next *Node
}
type Stack struct {
	top *Node
}

func main() {
	s := new(Stack)
	s.push(3)
	s.push(5)
	s.push(7)
	s.push(23)
	s.push(45)
	s.push(29)

	for !s.isEmpty() {
		fmt.Println(s.top.val)
		s.pop()
	}

}

func (s *Stack) push(x int) {
	newNode := &Node{val: x, next: s.top}
	s.top = newNode

}

func (s *Stack) pop() int {
	if s.top == nil {
		return 0
	}
	val := s.top.val
	s.top = s.top.next
	return val
}

func (s *Stack) topNode() int {
	if s.top == nil {
		return 0
	}
	return s.top.val
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}
