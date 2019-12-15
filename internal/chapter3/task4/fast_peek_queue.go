package task4

import "github.com/ivkalita/cti/internal/chapter3/common"

type fastPeekQueue struct {
	s1 *stack
	s2 *stack
}

func NewFastPeekQueue() *fastPeekQueue {
	return &fastPeekQueue{
		s1: &stack{},
		s2: &stack{},
	}
}

// Add – O(n) time, O(n) memory
func (q *fastPeekQueue) Add(val int) {
	for !q.s1.isEmpty() {
		el, _ := q.s1.pop()
		q.s2.push(el)
	}

	q.s2.push(val)

	for !q.s2.isEmpty() {
		el, _ := q.s2.pop()
		q.s1.push(el)
	}
}

// Remove – O(1) time, O(1) memory
func (q *fastPeekQueue) Remove() error {
	if q.s1.isEmpty() {
		return common.QueueIsEmpty
	}

	_, _ = q.s1.pop()

	return nil
}

// Peek – O(1) time, O(1) memory
func (q *fastPeekQueue) Peek() (int, error) {
	if q.s1.isEmpty() {
		return 0, common.QueueIsEmpty
	}

	return q.s1.top.data, nil
}

func (q *fastPeekQueue) IsEmpty() bool {
	return q.s1.isEmpty()
}
