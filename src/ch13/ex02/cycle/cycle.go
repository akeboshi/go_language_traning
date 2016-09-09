package cycle

import (
	"reflect"
	"unsafe"
)

func isCycle(v reflect.Value, seen []unsafe.Pointer) bool {
	if v.CanAddr() && v.Kind() != reflect.Array && v.Kind() != reflect.Slice && v.Kind() != reflect.Struct {
		vptr := unsafe.Pointer(v.UnsafeAddr())
		for _, p := range seen {
			if vptr == p {
				return true
			}
		}
		seen = append(seen, vptr)
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCycle(v.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if isCycle(v.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if isCycle(v.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range v.MapKeys() {
			if isCycle(v.MapIndex(k), seen) {
				return true
			}
		}
		return false
	}
	return false
}

func IsCycle(v interface{}) bool {
	seen := []unsafe.Pointer{}
	return isCycle(reflect.ValueOf(v), seen)
}
