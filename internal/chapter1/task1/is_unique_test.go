package task1_test

import (
	"github.com/ivkalita/cti/internal/chapter1/task1"
	"testing"
)

func TestIsUniqueWithSet(t *testing.T) {
	RunTests(t, task1.IsUniqueWithSet)
}

func TestIsUniqueWithSort(t *testing.T) {
	RunTests(t, task1.IsUniqueWithSort)
}
