package task3_test

import (
	"fmt"
	"github.com/ivkalita/cti/internal/chapter3/common"
	"github.com/ivkalita/cti/internal/chapter3/task3"
	"testing"
)

func TestStacksSet_Stack(t *testing.T) {
	maxStackSizes := []int{1, 2, 3, 4}
	for _, maxStackSize := range maxStackSizes {
		t.Run(fmt.Sprintf("max size %d", maxStackSize), func(t *testing.T) {
			factory := func() common.Stack {
				return task3.NewStacksSet(maxStackSize)
			}
			common.RunStackTests(t, factory)
		})
	}
}
