package task4

import "github.com/ivkalita/cti/internal/chapter3/common"

type fastAddQueue struct {
	s1 *stack
	s2 *stack
}

func NewFastAddQueue() *fastAddQueue {
	return &fastAddQueue{
		s1: &stack{},
		s2: &stack{},
	}
}

// Add – O(1) time, O(1) mem
func (q *fastAddQueue) Add(val int) {
	q.s1.push(val)
}

// Remove – O(n) time, O(n) mem
func (q *fastAddQueue) Remove() error {
	if q.s1.isEmpty() {
		return common.QueueIsEmpty
	}

	for !q.s1.isEmpty() {
		el, _ := q.s1.pop()
		q.s2.push(el)
	}

	_, _ = q.s2.pop()

	for !q.s2.isEmpty() {
		el, _ := q.s2.pop()
		q.s1.push(el)
	}

	return nil
}

// Peek – O(n) time, O(n) mem
func (q *fastAddQueue) Peek() (int, error) {
	if q.s1.isEmpty() {
		return 0, common.QueueIsEmpty
	}

	for !q.s1.isEmpty() {
		el, _ := q.s1.pop()
		q.s2.push(el)
	}

	r := q.s2.top.data

	for !q.s2.isEmpty() {
		el, _ := q.s2.pop()
		q.s1.push(el)
	}

	return r, nil
}

func (q *fastAddQueue) IsEmpty() bool {
	return q.s1.isEmpty()
}
