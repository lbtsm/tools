package slicehelper

import (
	"math"
)

/*
	包含以下功能
*/

// 按照size切割slice
func SplitSlice(src []string, size int) [][]string {
	if size < 1 {
		return nil
	}
	if len(src) <= size {
		return [][]string{src}
	}

	length := len(src)
	count := int(math.Ceil(float64(len(src)) / float64(size)))
	ret := make([][]string, 0, count)
	for i := 0; i < count; i++ {
		start := i * size
		end := start + size

		if start > length {
			start = length - 1
		}
		if end > length {
			end = length - 1
		}

		ret = append(ret, src[start:end:(size*(i+1))])
	}

	return ret
}

// 返回
func IdxOfStringSlice(slice []string, value string) int {
	if slice == nil {
		return -1
	}

	idx := -1
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return idx
}

// 是否存在在数组中
func ExistOfStringSlice(slice []string, value string) bool {
	if slice == nil {
		return false
	}

	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

// 转换为map
func ConvertToMap(slice []string) map[string]struct{} {
	if slice == nil || len(slice) == 0 {
		return nil
	}

	ret := make(map[string]struct{})
	for _, s := range slice {
		ret[s] = struct{}{}
	}

	return ret
}
