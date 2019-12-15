package task4_test

import (
	"github.com/ivkalita/cti/internal/chapter3/common"
	"github.com/ivkalita/cti/internal/chapter3/task4"
	"testing"
)

func TestFastAddQueue(t *testing.T) {
	factory := func() common.Queue {
		return task4.NewFastAddQueue()
	}

	common.RunQueueTests(t, factory)
}

func TestFastPeekQueue(t *testing.T) {
	factory := func() common.Queue {
		return task4.NewFastPeekQueue()
	}
	common.RunQueueTests(t, factory)
}
