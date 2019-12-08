package task3

import (
	"errors"
	"github.com/ivkalita/cti/internal/chapter3/common"
)

type StacksSet struct {
	stacks           []stack
	maxStackSize     int
	currentStackSize int
}

type stack struct {
	top *node
}

type node struct {
	data int
	next *node
}

func NewStacksSet(maxStackSize int) *StacksSet {
	return &StacksSet{
		stacks:           make([]stack, 0),
		maxStackSize:     maxStackSize,
		currentStackSize: maxStackSize,
	}
}

func (s *StacksSet) Push(item int) {
	if s.currentStackSize >= s.maxStackSize {
		s.stacks = append(s.stacks, stack{top: nil})
		s.currentStackSize = 0
	}
	s.currentStackSize++
	s.stacks[len(s.stacks)-1].Push(item)
}

func (s *StacksSet) Pop() (int, error) {
	if len(s.stacks) == 0 {
		return 0, common.StackIsEmpty
	}

	res, err := s.stacks[len(s.stacks)-1].Pop()
	if err != nil {
		return 0, errors.New("can't happen")
	}

	if s.currentStackSize == 1 {
		s.stacks = s.stacks[:len(s.stacks)-1]
		s.currentStackSize = s.maxStackSize + 1
	}
	s.currentStackSize--

	return res, nil
}

func (s StacksSet) IsEmpty() bool {
	return len(s.stacks) == 0
}

func (s *stack) Push(item int) {
	n := node{data: item}
	n.next = s.top
	s.top = &n
}

func (s *stack) Pop() (int, error) {
	if s.top == nil {
		return 0, common.StackIsEmpty
	}

	n := s.top
	s.top = s.top.next

	return n.data, nil
}
