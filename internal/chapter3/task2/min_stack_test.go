package task2_test

import (
	"github.com/ivkalita/cti/internal/chapter3/common"
	"github.com/ivkalita/cti/internal/chapter3/task2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStack_Stack(t *testing.T) {
	factory := func() common.Stack {
		return task2.NewMinStack()
	}

	common.RunStackTests(t, factory)
}

func TestMinStack_Min(t *testing.T) {
	tests := []struct {
		name     string
		executor func(t *testing.T, stack task2.MinStack)
	}{
		{
			name: "no elements -> error",
			executor: func(t *testing.T, stack task2.MinStack) {
				_, err := stack.Min()
				assert.Error(t, common.StackIsEmpty, err)
			},
		},
		{
			name: "one element minimum",
			executor: func(t *testing.T, stack task2.MinStack) {
				stack.Push(1)
				elem, err := stack.Min()
				assert.NoError(t, err)
				assert.Equal(t, 1, elem)
			},
		},
		{
			name: "two elements minimum",
			executor: func(t *testing.T, stack task2.MinStack) {
				stack.Push(1)
				stack.Push(0)
				elem, err := stack.Min()
				assert.NoError(t, err)
				assert.Equal(t, 0, elem)
			},
		},
		{
			name: "two elements, first is minimum",
			executor: func(t *testing.T, stack task2.MinStack) {
				stack.Push(1)
				stack.Push(2)
				elem, err := stack.Min()
				assert.NoError(t, err)
				assert.Equal(t, 1, elem)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stack := task2.NewMinStack()

			test.executor(t, *stack)
		})
	}
}
