package conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	var (
		i = 123
		s = "123"
	)
	assert.Equal(t, 123, Value(&i))
	assert.Equal(t, "123", Value(&s))
}

func TestValueOr(t *testing.T) {
	var i = 123
	assert.Equal(t, 123, ValueOr(&i, 123))
	assert.Equal(t, 321, ValueOr(nil, 321))
}

func TestPtrDefaultValue(t *testing.T) {
	testA := "123"
	assert.Equal(t, "123", PtrDefaultValue(&testA))
	var testB *string
	assert.Equal(t, "", PtrDefaultValue(testB))
}

func TestPtrValueOrNil(t *testing.T) {
	testA := "123"
	assert.Equal(t, &testA, PtrValueOrNil(testA))

	testB := 0
	assert.Equal(t, (*int)(nil), PtrValueOrNil(testB))
}
