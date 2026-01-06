package conv

import "golang.org/x/exp/constraints"

// Ptr 返回一个新的指针，指针指向的值与c相同
func Ptr[T any](t T) *T {
	return &t
}

// PtrDefaultValue 指针非空时对指针做浅拷贝，指针为空时返回对应类型的零值.
// Deprecated: 请使用 Value
func PtrDefaultValue[T any](p *T) T {
	if p == nil {
		var t T
		return t
	}
	return *p
}

// Value 指针非空时对指针指向的变量做浅拷贝，指针为空时返回对应类型的零值
func Value[T any](p *T) T {
	if p == nil {
		var t T
		return t
	}
	return *p
}

// ValueOr 类似于Value, 但是指针为空时返回传入的值而非类型零值
func ValueOr[T any](p *T, v T) T {
	if p != nil {
		return *p
	}
	return v
}

func PtrValueOrNil[T constraints.Ordered](t T) *T {
	// 如果t是该类型的零值，返回nil，否则返回t的指针
	var defaultT T
	if t == defaultT {
		return nil
	}
	return &t
}

func PtrValueOrDefault[T constraints.Ordered](t *T, defaultValue T) *T {
	// 如果t是nil，返回defaultValue，否则返回t的值
	if t == nil {
		return &defaultValue
	}
	return t
}

// DefaultValue 返回ptr指向的值，为空返回默认值
// Deprecated: 请使用 ValueOr
func DefaultValue[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

// EqualValue 判断两个指针指向的值是否相同
func EqualValue[T constraints.Ordered](t1 *T, t2 *T) bool {
	if t1 == nil {
		return t2 == nil
	}
	if t2 == nil {
		return false
	}
	return *t1 == *t2
}

// IntToIntPtr 将 T 类型的整数值转换为 K 类型的整数指针
func IntToIntPtr[T, K constraints.Integer](t T) *K {
	k := K(t)
	return &k
}

// IntPtrToInt 将 T 类型的整数指针转换为 K 类型的整数
func IntPtrToInt[T, K constraints.Integer](ptr *T, onNil K) K {
	if ptr == nil {
		return onNil
	}
	return K(*ptr)
}

// MapIntPtr 将 T 类型的整数指针转换为 K 类型的整数指针
func MapIntPtr[T, K constraints.Integer](ptr *T) *K {
	if ptr == nil {
		return nil
	}
	k := K(*ptr)
	return &k
}

// BoolToInt 将 bool 类型的转换为整数, true 为 1, false 为 0
func BoolToInt[T constraints.Integer](b bool) T {
	if b {
		return T(1)
	}
	return T(0)
}

func IntToBool[T constraints.Integer](i T) bool {
	return i > 0
}
