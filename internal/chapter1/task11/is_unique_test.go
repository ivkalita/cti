package task11_test

import (
	"github.com/ivkalita/cti/internal/chapter1/task11"
	"testing"
)

func TestIsUniqueWithSet(t *testing.T) {
	RunTests(t, task11.IsUniqueWithSet)
}

func TestIsUniqueWithSort(t *testing.T) {
	RunTests(t, task11.IsUniqueWithSort)
}
