package conv

import "golang.org/x/exp/constraints"

func NonDefaultOr[T comparable](judge, or T) T {
	var zero T
	if judge == zero {
		return or
	}

	return judge
}

// Int 将 T 类型的整数转换为 K 类型的整数
func Int[T, K constraints.Integer](t T) K {
	return K(t)
}

// IntSlice 将 T 类型的整数切片转换为 K 类型的整数切片
func IntSlice[T, K constraints.Integer](ts []T) []K {
	ks := make([]K, len(ts))
	for i, t := range ts {
		ks[i] = K(t)
	}
	return ks
}

func Must[T any](f func() (T, error)) T {
	v, err := f()
	if err != nil {
		panic(err)
	}
	return v
}
