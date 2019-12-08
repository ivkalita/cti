package common

import (
	"github.com/stretchr/testify/require"
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
				require.True(t, stack.IsEmpty())
			},
		},
		{
			name: "pushed not empty",
			executor: func(t *testing.T, stack Stack) {
				stack.Push(1)
				require.False(t, stack.IsEmpty())
			},
		},
		{
			name: "push/pop -> empty",
			executor: func(t *testing.T, stack Stack) {
				stack.Push(1)
				elem, err := stack.Pop()
				require.NoError(t, err)
				require.Equal(t, 1, elem)
				require.True(t, stack.IsEmpty())
			},
		},
		{
			name: "push push pop",
			executor: func(t *testing.T, stack Stack) {
				stack.Push(1)
				stack.Push(2)

				elem, err := stack.Pop()
				require.NoError(t, err)
				require.Equal(t, 2, elem)
			},
		},
		{
			name: "empty pop error",
			executor: func(t *testing.T, stack Stack) {
				_, err := stack.Pop()
				require.Equal(t, err, StackIsEmpty)
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
