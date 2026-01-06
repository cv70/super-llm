package conv

import (
	"golang.org/x/exp/constraints"
)

func IntAssertion[T constraints.Integer](value any) (T, bool) {
	switch v := value.(type) {
	case int:
		return T(v), true
	case int8:
		return T(v), true
	case int16:
		return T(v), true
	case int32:
		return T(v), true
	case int64:
		return T(v), true
	case uint:
		return T(v), true
	case uint8:
		return T(v), true
	case uint16:
		return T(v), true
	case uint32:
		return T(v), true
	case uint64:
		return T(v), true
	default:
		return 0, false
	}
}

func Assertion[T any](value any) T {
	var x T
	switch v := value.(type) {
	case T:
		return v
	default:
		return x
	}

}
