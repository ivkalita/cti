package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func RunStackTests(t *testing.T, newStack func() Stack) {
	tests := []struct {
		name     string
		executor func(t *testing.T, stack Stack)
	}{
		{
			name: "empty is empty",
			executor: func(t *testing.T, stack Stack) {
				assert.True(t, stack.IsEmpty())
			},
		},
		{
			name: "pushed not empty",
			executor: func(t *testing.T, stack Stack) {
				stack.Push(1)
				assert.False(t, stack.IsEmpty())
			},
		},
		{
			name: "push/pop -> empty",
			executor: func(t *testing.T, stack Stack) {
				stack.Push(1)
				elem, err := stack.Pop()
				assert.NoError(t, err)
				assert.Equal(t, 1, elem)
				assert.True(t, stack.IsEmpty())
			},
		},
		{
			name: "push push pop",
			executor: func(t *testing.T, stack Stack) {
				stack.Push(1)
				stack.Push(2)

				elem, err := stack.Pop()
				assert.NoError(t, err)
				assert.Equal(t, 2, elem)
			},
		},
		{
			name: "empty pop error",
			executor: func(t *testing.T, stack Stack) {
				_, err := stack.Pop()
				assert.Equal(t, StackIsEmpty, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := newStack()

			test.executor(t, s)
		})
	}
}

func RunQueueTests(t *testing.T, newQueue func() Queue) {
	tests := []struct {
		name     string
		executor func(t *testing.T, queue Queue)
	}{
		{
			name: "empty is empty",
			executor: func(t *testing.T, queue Queue) {
				assert.True(t, queue.IsEmpty())
			},
		},
		{
			name: "peek empty -> error",
			executor: func(t *testing.T, queue Queue) {
				_, err := queue.Peek()
				assert.Equal(t, QueueIsEmpty, err)
			},
		},
		{
			name: "peek non empty",
			executor: func(t *testing.T, queue Queue) {
				queue.Add(1)
				queue.Add(2)

				el, err := queue.Peek()

				assert.NoError(t, err)
				assert.Equal(t, 1, el)
			},
		},
		{
			name: "remove and add",
			executor: func(t *testing.T, queue Queue) {
				queue.Add(1)
				queue.Add(2)

				err := queue.Remove()
				assert.NoError(t, err)

				el, err := queue.Peek()
				assert.NoError(t, err)
				assert.Equal(t, 2, el)
			},
		},
		{
			name: "empty remove error",
			executor: func(t *testing.T, queue Queue) {
				err := queue.Remove()

				assert.Equal(t, QueueIsEmpty, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := newQueue()
			test.executor(t, q)
		})
	}
}
