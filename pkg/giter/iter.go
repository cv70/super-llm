package giter

import (
	"iter"
	"super-llm/pkg/gset"

	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Ordered](s iter.Seq[T]) T {
	var sum T
	for v := range s {
		sum += v
	}
	return sum
}

func Map[V1, V2 any](s iter.Seq[V1], mapFunc func(V1) V2) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		for item := range s {
			if !yield(mapFunc(item)) {
				return
			}
		}
	}
}

func Filter[V any](s iter.Seq[V], filterFunc func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for item := range s {
			if filterFunc(item) {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// Chunk 将迭代器按长度分片
func Chunk[V any](s iter.Seq[V], size int) iter.Seq[[]V] {
	if size <= 0 {
		panic("chunk size must be positive")
	}
	return func(yield func([]V) bool) {
		chunk := make([]V, 0, size)
		for item := range s {
			chunk = append(chunk, item)
			if len(chunk) == size {
				if !yield(chunk) {
					return
				}
				chunk = make([]V, 0, size)
			}
		}
		if len(chunk) > 0 {
			yield(chunk)
		}
	}
}

func Merge[T any](s ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range s {
			for item := range seq {
				if !yield(item) {
					return
				}
			}
		}
	}
}

func MergeDistinct[T comparable](s ...iter.Seq[T]) []T {
	seq := Merge(s...)
	set := gset.Set[T]{}
	for item := range seq {
		set.Add(item)
	}
	return set.ToSlice()
}

// Zip 合并两个迭代器，返回一个迭代器，每次迭代返回两个元素
func Zip[A any, B any](a iter.Seq[A], b iter.Seq[B]) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		aNext, aStop := iter.Pull(a)
		defer aStop()

		bNext, bStop := iter.Pull(b)
		defer bStop()

		for {
			aItem, aOk := aNext()
			bItem, bOk := bNext()
			if !aOk || !bOk || !yield(aItem, bItem) {
				return
			}
		}
	}
}

// ZipSeqSlice 合并一个迭代器和一个切片，返回一个迭代器，每次迭代返回两个元素
func ZipSeqSlice[A any, B any](b iter.Seq[B], a ...A) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		i := 0
		for bItem := range b {
			if i >= len(a) || !yield(a[i], bItem) {
				return
			}
		}
	}
}

// Seq2Slice 将迭代器转换为切片
func Seq2Slice[T any](seq iter.Seq[T]) []T {
	var result []T
	for item := range seq {
		result = append(result, item)
	}
	return result
}
