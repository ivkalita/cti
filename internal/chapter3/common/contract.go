package common

import "errors"

type Stack interface {
	Push(int)
	Pop() (int, error)
	IsEmpty() bool
}

var StackIsEmpty = errors.New("stack is empty")

type Queue interface {
	Add(int)
	Remove() error
	Peek() (int, error)
	IsEmpty() bool
}

var QueueIsEmpty = errors.New("queue is empty")
