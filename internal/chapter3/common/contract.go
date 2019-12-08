package common

import "errors"

type Stack interface {
	Push(int)
	Pop() (int, error)
	IsEmpty() bool
}

var StackIsEmpty = errors.New("stack is empty")
