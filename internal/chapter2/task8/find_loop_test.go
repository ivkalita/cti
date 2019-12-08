package task8_test

import (
	"github.com/ivkalita/cti/internal/chapter2/task8"
	"testing"
)

func TestFindLoop(t *testing.T) {
	RunTests(t, task8.FindLoop)
}
