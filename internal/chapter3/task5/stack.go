package task5

import "github.com/ivkalita/cti/internal/chapter3/common"

type stack struct {
	top *node
}

func NewStack() common.Stack {
	return &stack{}
}

func (s *stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, common.StackIsEmpty
	}

	r := s.top.data

	s.top = s.top.next

	return r, nil
}

func (s *stack) IsEmpty() bool {
	return s.top == nil
}

func (s *stack) Push(a int) {
	s.top = &node{data: a, next: s.top}
}

type node struct {
	data int
	next *node
}
