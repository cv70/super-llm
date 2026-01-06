package gset

import "iter"

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Clear() {
	clear(s)
}

func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, s.Len())
	for k := range s {
		slice = append(slice, k)
	}
	return slice
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	union := make(Set[T])
	for k := range s {
		union.Add(k)
	}
	for k := range other {
		union.Add(k)
	}
	return union
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	intersect := make(Set[T])
	for k := range s {
		if other.Has(k) {
			intersect.Add(k)
		}
	}
	return intersect
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	difference := make(Set[T])
	for k := range s {
		if !other.Has(k) {
			difference.Add(k)
		}
	}
	return difference
}

func (s Set[T]) In(other Set[T]) bool {
	for k := range s {
		if other.Has(k) {
			return true
		}
	}
	return false
}

func (s Set[T]) Range() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for k := range s {
			if !yield(i, k) {
				return
			}
			i++
		}
	}
}
