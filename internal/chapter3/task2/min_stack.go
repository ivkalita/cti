package task2

import (
	"github.com/ivkalita/cti/internal/chapter3/common"
)

type MinStack struct {
	top *node
}

type node struct {
	data int
	min  int
	next *node
}

func NewMinStack() *MinStack {
	return &MinStack{}
}

func (ms *MinStack) Push(item int) {
	newMin := item
	if ms.top != nil && ms.top.min < newMin {
		newMin = ms.top.min
	}

	n := node{data: item, min: newMin}
	n.next = ms.top
	ms.top = &n
}

func (ms *MinStack) Pop() (int, error) {
	if ms.top == nil {
		return 0, common.StackIsEmpty
	}

	res := ms.top.data
	ms.top = ms.top.next

	return res, nil
}

func (ms MinStack) Min() (int, error) {
	if ms.top == nil {
		return 0, common.StackIsEmpty
	}

	return ms.top.min, nil
}

func (ms MinStack) IsEmpty() bool {
	return ms.top == nil
}
