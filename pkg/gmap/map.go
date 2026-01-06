package gmap

import (
    "encoding/json"
    "fmt"
    "iter"
    "math"
    "slices"
    "sync"

    "golang.org/x/exp/constraints"
)

// Reverse 把一个 map 的 key 和 value 转换，返回一个新的map
func Reverse[K comparable, V comparable](m map[K]V) map[V]K {
    r := make(map[V]K, len(m))
    for k, v := range m {
        r[v] = k
    }
    return r
}

// Diff 比较 base 和 another，返回 another 中的某些值
// 1. 不在 base 里的
// 2. 在 base 里，但是和 base 里的 value 不一样的)
func Diff[K comparable, V comparable](base map[K]V, another map[K]V) map[K]V {
    ret := make(map[K]V)
    for k, v := range another {
        v2, ok := base[k]
        if (!ok) || (v2 != v) {
            ret[k] = v
        }
    }
    return ret
}

func Get[K comparable, V any](m map[K]V, key K, defaultVal V) V {
    if v, ok := m[key]; ok {
        return v
    }
    return defaultVal
}

func Items[K comparable, V any, T any](m map[K]V, itemFunc func(K, V) T) []T {
    items := make([]T, 0, len(m))
    for key, value := range m {
		item := itemFunc(key, value)
		items = append(items, item)
    }
    return items
}

func ItemsIf[K comparable, V any, T any](m map[K]V, itemFunc func(K, V) (T, bool)) []T {
    items := make([]T, 0, len(m))
    for key, value := range m {
		item, ok := itemFunc(key, value)
		if ok {
			items = append(items, item)
		}
    }
    return items
}

func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    return keys
}

func KeysIf[K comparable, V any](m map[K]V, f func(K, V) bool) []K {
    keys := make([]K, 0, len(m))
    for key, value := range m {
        if f(key, value) {
            keys = append(keys, key)
        }
    }
    return keys
}

func KeysContainsAny[K comparable, V any](m map[K]V, arr ...K) bool {
    for _, val := range arr {
        if _, ok := m[val]; ok {
            return true
        }
    }
    return false
}

func AppendWithEmptyValue[K comparable](m map[K]struct{}, arr ...K) map[K]struct{} {
    if m == nil {
        m = make(map[K]struct{}, len(arr))
    }
    for _, val := range arr {
        m[val] = struct{}{}
    }
    return m
}

func AnyMatchIf[K comparable, V any](m map[K]V, f func(K, V) bool) bool {
    for k, v := range m {
        if f(k, v) {
            return true
        }
    }
    return false
}

func Values[K comparable, V any](m map[K]V) []V {
    values := make([]V, 0, len(m))
    for _, value := range m {
        values = append(values, value)
    }
    return values
}

func ValuesSum[K comparable, V constraints.Ordered](m map[K]V) V {
    var sum V
    for _, value := range m {
        sum += value
    }
    return sum
}

func ValuesMap[K comparable, V1 any, V2 any](m map[K]V1, f func(K, V1) V2) []V2 {
    values := make([]V2, 0, len(m))
    for key, value := range m {
        values = append(values, f(key, value))
    }
    return values
}

func ValuesIf[K comparable, V any](m map[K]V, f func(K, V) bool) []V {
    values := make([]V, 0, len(m))
    for key, value := range m {
        if f(key, value) {
            values = append(values, value)
        }
    }
    return values
}

func Serialize[K comparable, V1 any, V2 any](m map[K]V1, serializeFunc func(V1) V2) map[K]V2 {
    res := make(map[K]V2, len(m))
    for k, v := range m {
        res[k] = serializeFunc(v)
    }
    return res
}

func SerializeIf[K comparable, V1 any, V2 any](m map[K]V1, serializeFunc func(V1) (V2, bool)) map[K]V2 {
    res := make(map[K]V2)
    for k, v := range m {
        if _v, ok := serializeFunc(v); ok {
            res[k] = _v
        }
    }
    return res
}

func GetFields[K comparable, V1 any, V2 any](m map[K]V1, getField func(V1) V2) []V2 {
    res := make([]V2, 0, len(m))
    for _, v := range m {
        res = append(res, getField(v))
    }
    return res
}

// Merge is deprecated
// Deprecated: use maps.Copy after go 1.21
func Merge[K comparable, V any](m1, m2 map[K]V) map[K]V {
    res := make(map[K]V, len(m1)+len(m2))
    for k, v := range m1 {
        res[k] = v
    }
    for k, v := range m2 {
        res[k] = v
    }
    return res
}

// SyncMapToMap 把 sync.Map 转换成 map，如果类型不匹配，会忽略
func SyncMapToMap[K comparable, V any](m *sync.Map) map[K]V {
    res := make(map[K]V)
    m.Range(func(key, value any) bool {
        k, ok1 := key.(K)
        v, ok2 := value.(V)
        if ok1 && ok2 {
            res[k] = v
        }
        return true
    })
    return res
}

func SubMap[K comparable, V any](m map[K]V, keys []K) map[K]V {
    res := make(map[K]V, len(keys))
    for _, k := range keys {
        if v, ok := m[k]; ok {
            res[k] = v
        }
    }
    return res
}

func Chunked[K comparable, V any](m map[K]V, size int) []map[K]V {
    mapCount := int(math.Ceil(float64(len(m)) / float64(size)))
    res := make([]map[K]V, 0, mapCount)
    tmpMap := make(map[K]V, size)
    now := 0
    for k, v := range m {
        now += 1
        tmpMap[k] = v
        if now == size {
            now = 0
            res = append(res, tmpMap)
            tmpMap = make(map[K]V, size)
        }
    }
    if len(tmpMap) != 0 {
        res = append(res, tmpMap)
    }
    return res
}

func HasKey[M ~map[K]V, K comparable, V any](m M, key K) bool {
    _, ok := m[key]
    return ok
}

func SliceToMap[K comparable, V any](s []V, keyFunc func(V) K) map[K]V {
    res := make(map[K]V, len(s))
    for _, v := range s {
        res[keyFunc(v)] = v
    }
    return res
}

func SliceToKVMapIf[K comparable, T, V any](s []T, kvFunc func(T) (K, V, bool)) map[K]V {
    res := make(map[K]V, len(s))
    for _, t := range s {
        if k, v, ok := kvFunc(t); ok {
            res[k] = v
        }
    }
    return res
}

func SliceToKVMap[K comparable, T, V any](s []T, kvFunc func(T) (K, V)) map[K]V {
    res := make(map[K]V, len(s))
    for _, t := range s {
        k, v := kvFunc(t)
        res[k] = v
    }
    return res
}

func SliceToSliceMap[K comparable, V any](s []V, keyFunc func(V) K) map[K][]V {
    res := make(map[K][]V, len(s))
    for _, v := range s {
        res[keyFunc(v)] = append(res[keyFunc(v)], v)
    }
    return res
}

func SliceToSliceKVMap[K, V comparable, T any](s []T, kvFunc func(T) (K, V)) map[K][]V {
    res := make(map[K][]V, len(s))
    for _, t := range s {
        k, v := kvFunc(t)
        res[k] = append(res[k], v)
    }
    return res
}

func CompareMapEqual[K comparable, V comparable](map1, map2 map[K]V) bool {
    if len(map1) != len(map2) {
        return false
    }
    for k, map1V := range map1 {
        if map2V, ok := map2[k]; !ok {
            return false
        } else {
            if map1V != map2V {
                return false
            }
        }
    }
    return true
}

// GetOrderedMapStr 对Map排序并输出string
func GetOrderedMapStr[K comparable, V any](m map[K]V, less func(a, b V) int) string {
    type Pair[T1 any, T2 any] struct {
        First  T1
        Second T2
    }

    var ret []Pair[K, V]
    for k, v := range m {
        ret = append(ret, Pair[K, V]{k, v})
    }
    slices.SortFunc(ret, func(a, b Pair[K, V]) int {
        return less(a.Second, b.Second)
    })
    return fmt.Sprintf("%v", ret)
}

func ChunkMap[k comparable, v any](m map[k]v, chunkSize int) []map[k]v {
    mapLen := (len(m) + chunkSize - 1) / chunkSize
    res := make([]map[k]v, 0, mapLen)

    cnt := 0
    for key, value := range m {
        if cnt%chunkSize == 0 {
            if len(res) == mapLen-1 {
                res = append(res, make(map[k]v, len(m)%chunkSize))
            } else {
                res = append(res, make(map[k]v, chunkSize))
            }
        }
        res[len(res)-1][key] = value
        cnt++
    }
    return res
}

func MapValueMap[K comparable, V, T any](m map[K]V, valueFunc func(V) T) map[K]T {
    res := make(map[K]T, len(m))
    for key, value := range m {
        newValue := valueFunc(value)
        res[key] = newValue
    }
    return res
}

func MapValueMapIf[K comparable, V, T any](m map[K]V, valueFunc func(V) (T, bool)) map[K]T {
    res := make(map[K]T, len(m))
    for key, value := range m {
        newValue, ok := valueFunc(value)
        if ok {
            res[key] = newValue
        }
    }
    return res
}

func MapKeyMap[K, T comparable, V any](m map[K]V, keyFunc func(K) T) map[T]V {
    res := make(map[T]V, len(m))
    for key, value := range m {
        newKey := keyFunc(key)
        res[newKey] = value
    }
    return res
}

func MapKeyMapIf[K, T comparable, V any](m map[K]V, keyFunc func(K) (T, bool)) map[T]V {
    res := make(map[T]V, len(m))
    for key, value := range m {
        newKey, ok := keyFunc(key)
		if ok {
			res[newKey] = value
		}
    }
    return res
}

func PutIfAbsent[K comparable, V any](m map[K]V, k K, v V) bool {
    _, exists := m[k]
    if !exists {
        m[k] = v
        return true
    }
    return false
}

// SingleMap 一个kv的map
func SingleMap[K comparable, V any](k K, v V) map[K]V {
    result := make(map[K]V, 1)
    result[k] = v
    return result
}

// StructToMap struct转map
func StructToMap[K comparable, V any, T any](entity T) (map[K]V, error) {
    // 将结构体转成 JSON 字符串
    jsonData, err := json.Marshal(entity)
    if err != nil {
        return nil, err
    }

    // 将 JSON 字符串转换为 map[K]V
    var result map[K]V
    err = json.Unmarshal(jsonData, &result)
    if err != nil {
        return nil, err
    }

    return result, nil
}

func SliceToSliceMapWithKV[K comparable, V any, T any](s []V, kvFunc func(V) (K, T)) map[K]T {
    res := make(map[K]T, len(s))
    for _, v := range s {
        k, t := kvFunc(v)
        res[k] = t
    }
    return res
}

func Chain[K comparable, V any](maps ...map[K]V) iter.Seq2[K, V] {
    return func(yield func(K, V) bool) {
        for _, m := range maps {
            for k, v := range m {
                if !yield(k, v) {
                    return
                }
            }
        }
    }
}
