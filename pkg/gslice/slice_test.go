package gslice

import (
	"testing"
)

func TestRangeNilSlice(t *testing.T) {
	var s []int = nil
	for _, item := range s {
		t.Log(item)
	}
}
