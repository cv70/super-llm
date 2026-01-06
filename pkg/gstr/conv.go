package gstr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"unsafe"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

// StringToBytes 把字符串转换成字节切片(无需内存复制)
// 参考: https://github.com/gin-gonic/gin/blob/a481ee2897af1e368de5c919fbeb21b89aa26fc7/internal/bytesconv/bytesconv_1.20.go
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// BytesToString 把字节序列转换成字符串(无需内存复制)
func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func BoolToString[T bool](a T) string {
	return strconv.FormatBool(bool(a))
}

func IntToString[T constraints.Signed](a T) string {
	return strconv.FormatInt(int64(a), 10)
}

func FloatToString[T constraints.Float](a T) string {
	return strconv.FormatFloat(float64(a), 'f', -1, 64)
}

func UIntToString[T constraints.Unsigned](a T) string {
	return strconv.FormatUint(uint64(a), 10)
}

func StructToString(someThing any) string {
	b, err := json.Marshal(someThing)
	if err != nil {
		return fmt.Sprintf("%+v", someThing)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", someThing)
	}
	return out.String()
}

func StringToInt64Array(str string, split string) []int64 {
	if str == "" {
		return []int64{}
	}
	strList := strings.Split(str, split)
	resp := make([]int64, 0, len(strList))
	for i := 0; i < len(strList); i++ {
		v, err := strconv.ParseInt(strList[i], 10, 64)
		if err != nil {
			continue
		}
		resp = append(resp, v)
	}
	return resp
}

func StringToStrArray(str string, split string) []string {
	if str == "" {
		return []string{}
	}
	strList := strings.Split(str, split)
	resp := make([]string, 0, len(strList))
	for i := 0; i < len(strList); i++ {
		if strList[i] != "" {
			resp = append(resp, strList[i])
		}
	}
	return resp
}

func StringToInt32Array(str string, split string) []int32 {
	if str == "" {
		return []int32{}
	}
	strList := strings.Split(str, split)
	resp := make([]int32, 0, len(strList))
	for i := 0; i < len(strList); i++ {
		v, err := strconv.ParseInt(strList[i], 10, 64)
		if err != nil {
			continue
		}
		resp = append(resp, int32(v))
	}
	return resp
}

func IntJoinToString[T int | int32 | int64](s []T, sep string) string {
	strSlice := make([]string, len(s))

	for i, v := range s {
		strSlice[i] = strconv.FormatInt(int64(v), 10)
	}

	return strings.Join(strSlice, sep)
}

func StringSliceToIntSlice[T constraints.Signed](s []string) ([]T, error) {
	intSlice := make([]T, len(s))

	for i, v := range s {
		intValue, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to convert string to int: %s", v)
		}
		intSlice[i] = T(intValue)
	}

	return intSlice, nil
}
