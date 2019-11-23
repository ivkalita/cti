package task3_test

import (
	"github.com/ivkalita/cti/internal/chapter2/task3"
	"testing"
)

func TestRemoveNode(t *testing.T) {
	RunTests(t, task3.RemoveNode)
}
