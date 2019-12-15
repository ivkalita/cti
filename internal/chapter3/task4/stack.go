package task4

import "github.com/ivkalita/cti/internal/chapter3/common"

type stack struct {
	top *node
}

func (s *stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, common.StackIsEmpty
	}

	r := s.top.data

	s.top = s.top.next

	return r, nil
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

func (s *stack) push(a int) {
	s.top = &node{data: a, next: s.top}
}

type node struct {
	data int
	next *node
}
