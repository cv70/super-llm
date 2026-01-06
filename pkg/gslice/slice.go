package gslice

import (
	"fmt"
	"iter"
	"math/rand"
	"slices"
	"strings"
	"super-llm/pkg/gset"
	"time"
)

func ContainsAny[T comparable](s []T, values ...T) bool {
	return HasIntersection(s, values)
}

func HasIntersection[T comparable](a, b []T) bool {
	if (len(a) == 0) || (len(b) == 0) {
		return false
	}
	if len(a) > len(b) {
		a, b = b, a
	}
	set := make(map[T]struct{})

	for i := 0; i < len(a); i++ {
		set[a[i]] = struct{}{}
	}
	for i := 0; i < len(b); i++ {
		if _, ok := set[b[i]]; ok {
			return true
		}
	}

	return false
}

// HasIntersectionFor 使用for循环实现的 HasIntersection
func HasIntersectionFor[T comparable](a, b []T) bool {
	if (len(a) == 0) || (len(b) == 0) {
		return false
	}

	for _, itemA := range a {
		for _, itemB := range b {
			if itemA == itemB {
				return true
			}
		}
	}
	return false
}

func Intersection[T comparable](a, b []T) []T {
	if (len(a) == 0) || (len(b) == 0) {
		return nil
	}
	if len(a) > len(b) {
		a, b = b, a
	}

	set := make(map[T]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}

	var ret []T
	for _, v := range b {
		if _, ok := set[v]; ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func RemoveDups[T comparable](s []T) []T {
	return RemoveDupsFunc(s, func(item T) T {
		return item
	})
}

func RemoveDupsFunc[T any, K comparable](s []T, getKey func(T) K) (ret []T) {
	set := make(map[K]struct{})
	for i := 0; i < len(s); i++ {
		key := getKey(s[i])
		if _, ok := set[key]; !ok {
			set[key] = struct{}{}
			ret = append(ret, s[i])
		}
	}
	return ret
}

func SliceToAny[T any](s []T) []any {
	res := make([]any, 0, len(s))
	for _, item := range s {
		res = append(res, item)
	}
	return res
}

func GetPart[T any](s []T, offset int, limit int) []T {
	if len(s) <= offset {
		return nil
	}
	res := s[offset:]
	if len(res) < limit {
		return res
	}
	return res[:limit]
}

func GetPartLimit[T any](s []T, limit int) []T {
	return GetPart(s, 0, limit)
}

func AnyToString(s []any) []string {
	res := make([]string, 0, len(s))
	for _, item := range s {
		res = append(res, fmt.Sprintf("%v", item))
	}
	return res
}

func SliceToSet[T comparable](s []T) map[T]bool {
	m := make(map[T]bool, len(s))
	for _, item := range s {
		m[item] = true
	}
	return m
}

func SliceToSetStruct[T comparable](s []T) map[T]struct{} {
	m := make(map[T]struct{}, len(s))
	for _, item := range s {
		m[item] = struct{}{}
	}
	return m
}

func Serialize[T1 any, T2 any](src []T1, serializeFunc func(T1) T2) []T2 {
	dst := make([]T2, len(src))
	for i := range src {
		dst[i] = serializeFunc(src[i])
	}
	return dst
}

func SliceToMap[T comparable, V any](s []V, getKey func(V) T) map[T]V {
	res := make(map[T]V, len(s))
	for _, item := range s {
		key := getKey(item)
		res[key] = item
	}
	return res
}

func SliceToKVMap[K comparable, T, V any](s []T, getKV func(T) (K, V)) map[K]V {
	res := make(map[K]V, len(s))
	for _, item := range s {
		key, value := getKV(item)
		res[key] = value
	}
	return res
}

func SliceToKVMapIf[K comparable, T, V any](s []T, getKV func(T) (K, V, bool)) map[K]V {
	res := make(map[K]V, len(s))
	for _, item := range s {
		key, value, ok := getKV(item)
		if ok {
			res[key] = value
		}
	}
	return res
}

func Slice2Set[T comparable](s []T) gset.Set[T] {
	set := gset.Set[T]{}
	for _, item := range s {
		set.Add(item)
	}
	return set
}

func GetFields[V1 any, V2 any](s []V1, getField func(V1) V2) []V2 {
	res := make([]V2, 0, len(s))
	for _, item := range s {
		res = append(res, getField(item))
	}
	return res
}

func GetFieldsWithFilter[V1 any, V2 any](s []V1, getField func(V1) (V2, bool)) []V2 {
	res := make([]V2, 0, len(s))
	for _, item := range s {
		if newItem, ok := getField(item); ok {
			res = append(res, newItem)
		}
	}
	return res
}

func GetFieldsFirstWithFilter[V1 any, V2 any](s []V1, getField func(V1) (V2, bool)) V2 {
	var res V2
	for _, item := range s {
		if newItem, ok := getField(item); ok {
			res = newItem
			break
		}
	}
	return res
}

func TopN[T any](s []T, n int) []T {
	return s[:min(len(s), n)]
}

func LastN[T any](s []T, n int) []T {
	return s[max(0, len(s)-n):]
}

// Slice 直接在原切片上做slice
func Slice[T any](s []T, start, end int) []T {
	_start := min(max(0, start), len(s))
	_end := max(min(len(s), end), _start)
	return s[_start:_end]
}

// DiffSliceSingle (slice1 - slice2)
func DiffSliceSingle[T comparable](slice1 []T, slice2 []T) []T {
	inter := make([]T, 0, len(slice1))
	mp := make(map[T]bool)
	for _, s := range slice2 {
		mp[s] = true
	}
	for _, s := range slice1 {
		if _, ok := mp[s]; ok {
			continue
		}
		inter = append(inter, s)
	}
	return inter
}

// DiffFunc (slice1 - slice2)
func DiffFunc[T any, K comparable](slice1 []T, slice2 []T, getKey func(T) K) []T {
	inter := make([]T, 0, len(slice1))
	mp := make(map[K]struct{})
	for _, s := range slice2 {
		mp[getKey(s)] = struct{}{}
	}
	for _, s := range slice1 {
		if _, ok := mp[getKey(s)]; ok {
			continue
		}
		inter = append(inter, s)
	}
	return inter
}

// Compare get slice diff data
// less = slice2 - slice1
// more = slice1 - slice2
func Compare[T comparable](slice1 []T, slice2 []T) (less []T, more []T, equal []T) {
	m1 := make(map[T]struct{})
	for _, s := range slice1 {
		m1[s] = struct{}{}
	}
	m2 := make(map[T]struct{})
	for _, s := range slice2 {
		m2[s] = struct{}{}
	}
	equalSet := make(map[T]struct{})
	for k := range m1 {
		if _, ok := m2[k]; !ok {
			more = append(more, k)
		} else {
			equalSet[k] = struct{}{}
		}
	}
	for k := range m2 {
		if _, ok := m1[k]; !ok {
			less = append(less, k)
		} else {
			equalSet[k] = struct{}{}
		}
	}
	for k := range equalSet {
		equal = append(equal, k)
	}
	return less, more, equal
}

func GetRef[T any](s []T, idx int) *T {
	length := len(s)
	if idx < 0 {
		idx += length
	}
	if idx < 0 || idx >= length {
		panic("index out of range")
	}
	return &s[idx]
}

// TryGet 类似于Get, 但是当下标越界时会返回类型的零值.
func TryGet[S ~[]T, T any](s S, idx int) T {
	length := len(s)
	if idx < 0 {
		idx += length
	}
	if idx < 0 || idx >= length {
		var t T
		return t
	}
	return s[idx]
}

func Reduce[T any, V any](s []T, reduceFunc func(V, T) V) V {
	var init V
	for _, item := range s {
		init = reduceFunc(init, item)
	}
	return init
}

func Merge[T any](s ...[]T) []T {
	totalLen := Reduce(s, func(acc int, item []T) int {
		return acc + len(item)
	})
	res := make([]T, 0, totalLen)
	for _, val := range s {
		res = append(res, val...)
	}
	return res
}

func MergeDistinct[T comparable](ss ...[]T) []T {
	var (
		ret []T
		set = make(map[T]struct{})
	)
	for _, s := range ss {
		for _, v := range s {
			if _, ok := set[v]; !ok {
				set[v] = struct{}{}
				ret = append(ret, v)
			}
		}
	}
	return ret
}

func MergerInOrder[T any](s ...[]T) []T {
	var ret []T
	for now := 0; ; now++ {
		flag := false
		for _, v := range s {
			if now >= len(v) {
				continue
			}
			ret = append(ret, v[now])
			flag = true
		}
		if !flag {
			break
		}
	}
	return ret
}

// RemoveVal 从切片中移除所有值为target的元素，返回新的切片
func RemoveVal[T comparable](s []T, target T) []T {
	return Filter(s, func(v T) bool {
		return v != target
	})
}

// MoveElementsToFront 将满足要求的元素移到列表的前面, 并返回一个使得s[:i]为全部满足要求元素的下标i.
func MoveElementsToFront[T any](s []T, check func(T) bool) int {
	var l int
	for r := len(s) - 1; l != r; {
		if check(s[r]) {
			s[l], s[r] = s[r], s[l]
			l++
		} else {
			r--
		}
	}

	if check(s[l]) {
		return l + 1
	}
	return l
}

// MoveElementsToFrontInOrder 将一个数组中check为true的元素移动至数组开头，被移动元素的相对位置不变
// 同时保持数组A中其他未被移动元素的相对位置不变，时间复杂度为O(N)。
func MoveElementsToFrontInOrder[T any](s []T, check func(T) bool) []T {
	var (
		tmpElements []T
		l, r        int
	)
	for l, r = len(s)-1, len(s)-1; l >= 0; {
		s[r] = s[l]
		if check(s[l]) {
			tmpElements = append(tmpElements, s[l])
		} else {
			r--
		}
		l--
	}
	for _, item := range tmpElements {
		s[r] = item
		r--
	}
	return s
}

func Any[T any](s []T, check func(T) bool) bool {
	for _, v := range s {
		if check(v) {
			return true
		}
	}
	return false
}

func All[T any](s []T, check func(T) bool) bool {
	for _, v := range s {
		if !check(v) {
			return false
		}
	}
	return true
}

// GetElementAtPercentage 获取有序数组中第百分之x的内容
func GetElementAtPercentage[T any](items []T, percentage float64) T {
	if len(items) == 0 {
		var t T
		return t
	}
	idx := int(float64(len(items)) * percentage)
	if idx < 0 {
		idx = 0
	} else if idx >= len(items) {
		idx = len(items) - 1
	}
	return items[idx]
}

// Shuffle  随机打乱数组
func Shuffle[T any](arr []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)              // 生成一个 [0, i] 之间的随机整数
		arr[i], arr[j] = arr[j], arr[i] // 交换 arr[i] 和 arr[j] 的值
	}
}

// RandGet 随机获取数组中的n个元素
func RandGet[T any](arr []T, n int) []T {
	Shuffle(arr)
	if n >= len(arr) {
		return arr
	}
	return arr[:n]
}

// SliceTransform 遍历slice的值并转换
func SliceTransform[T comparable, V any](s []V, getValue func(V) T) []T {
	result := make([]T, 0, len(s))
	for _, item := range s {
		value := getValue(item)
		result = append(result, value)
	}
	return result
}

// SliceLogicTransform  遍历slice的值并转换
func SliceLogicTransform[T comparable, V any](s []V, getValue func(V) (T, bool)) []T {
	result := make([]T, 0, len(s))
	for _, item := range s {
		value, ok := getValue(item)
		if ok {
			result = append(result, value)
		}
	}
	return result
}

// SingleSlice 构建一个元素的切片
func SingleSlice[T comparable](item T) []T {
	result := make([]T, 0, 1)
	result = append(result, item)
	return result
}

// EmptySlice 构建无元素切片
func EmptySlice[T comparable]() []T {
	result := make([]T, 0)
	return result
}

// FirstMFromEveryN containing the first M elements from every group of N elements.
func FirstMFromEveryN[T any](items []T, m int, n int) []T {
	var result []T
	for i := 0; i < len(items); i += n {
		for j := i; j < i+m && j < len(items); j++ {
			result = append(result, items[j])
		}
	}
	return result
}

func FirstIf[T any](items []T, check func(T) bool) (T, int) {
	for i := range len(items) {
		if check(items[i]) {
			return items[i], i
		}
	}
	var zero T
	return zero, -1
}

func LastIf[T any](items []T, check func(T) bool) (T, int) {
	for i := len(items) - 1; i >= 0; i-- {
		if check(items[i]) {
			return items[i], i
		}
	}
	var zero T
	return zero, -1
}

func SplitAndTrimSpace(s string, sep string) []string {
	ret := strings.Split(s, sep)
	for i := range ret {
		ret[i] = strings.TrimSpace(ret[i])
	}
	return ret
}

func ToValueIndexMap[T comparable](s []T) map[T]int {
	ret := make(map[T]int, len(s))
	for i, val := range s {
		ret[val] = i
	}
	return ret
}

func CountFunc[T any](s []T, check func(T) bool) int {
	var count int
	for _, e := range s {
		if check(e) {
			count++
		}
	}
	return count
}

func Count[T comparable](s []T, target T) int {
	return CountFunc(s, func(e T) bool {
		return e == target
	})
}

func Chain[T comparable](s ...[]T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s {
			for _, item := range v {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// PercentileFunc 获取一个数组的分位元素，使用排序实现，要求数组长度大于0，会修改原数组
func PercentileFunc[T any](arr []T, percent float64, cmp func(a, b T) int) T {
	if len(arr) == 0 {
		panic("array length must be greater than 0")
	}
	if percent < 0 || percent > 1 {
		panic("percent must be between 0 and 1")
	}
	slices.SortFunc(arr, cmp)
	index := max(int(float64(len(arr))*percent)-1, 0)
	return arr[index]
}

// ZipSlice 合并两个切片，返回一个迭代器，每次迭代返回两个元素
func ZipSlice[A any, B any](a []A, b []B) iter.Seq2[A, B] {
	length := min(len(a), len(b))
	return func(yield func(A, B) bool) {
		for i := range length {
			if !yield(a[i], b[i]) {
				return
			}
		}
	}
}

// Slice2Seq 将切片转换为迭代器
func Slice2Seq[T any](slice ...T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range slice {
			if !yield(item) {
				return
			}
		}
	}
}
