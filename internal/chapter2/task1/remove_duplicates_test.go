package task1_test

import (
	"github.com/ivkalita/cti/internal/chapter2/task1"
	"testing"
)

func TestRemoveDuplicatesWithMap(t *testing.T) {
	RunTests(t, task1.RemoveDuplicatesWithMap)
}

func TestRemoveDuplicatesWithoutMap(t *testing.T) {
	RunTests(t, task1.RemoveDuplicatesWithoutMap)
}
