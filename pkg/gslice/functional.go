package gslice

func Filter[S ~[]T, T any](items S, check func(T) bool) S {
	var ret S
	for _, e := range items {
		if check(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

// FilterInPlace 原地过滤掉check=false的元素
func FilterInPlace[S ~[]T, T any](items *S, check func(T) bool) {
	var j int
	for _, e := range *items {
		if check(e) {
			(*items)[j] = e
			j++
		}
	}
	*items = (*items)[:j]
}

func Map[T, U any](items []T, f func(T) U) []U {
	ret := make([]U, len(items))
	for i, e := range items {
		ret[i] = f(e)
	}
	return ret
}

func MapIf[T, U any](items []T, f func(T) (U, bool)) []U {
	var ret []U
	for _, e := range items {
		if u, ok := f(e); ok {
			ret = append(ret, u)
		}
	}
	return ret
}

func MapExtend[T, U any](items []T, f func(T) []U) []U {
	ret := make([]U, 0, len(items))
	for _, e := range items {
		ret = append(ret, f(e)...)
	}
	return ret
}

func Accumulate[T, U any](items []U, init T, addFunc func(T, U) T) T {
	for _, e := range items {
		init = addFunc(init, e)
	}
	return init
}

func Partition[S ~[]T, T any](items S, check func(T) bool) (hit S, miss S) {
	for _, e := range items {
		if check(e) {
			hit = append(hit, e)
		} else {
			miss = append(miss, e)
		}
	}
	return hit, miss
}
