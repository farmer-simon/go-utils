package egu

import (
	"fmt"
	"reflect"
	"strings"
)

// Implode concat items by the given separator
func Implode( items interface{}, separator string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(items), "[]"), " ", separator, -1)
}

// Explode split string with separator
func Explode(str, separator string) StringSlice {
	return String(strings.Split(str, separator))
}

// InArray
// haystack supported types: slice, array or map
func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type must be slice, array or map")
	}

	return false
}
