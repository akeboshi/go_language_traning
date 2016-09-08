// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package format

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Slice, reflect.Array:
		arr := []string{}
		for i := 0; i < v.Len(); i++ {
			arr = append(arr, formatAtom(v.Index(i)))
		}
		return v.Type().String() + "{" + strings.Join(arr, ", ") + "}"
	case reflect.Struct:
		m := [][2]string{}
		for i := 0; i < v.NumField(); i++ {
			a := formatAtom(v.Field(i))
			m = append(m, [2]string{v.Type().Field(i).Name, a})
		}
		sArr := []string{}
		for _, v := range m {
			sArr = append(sArr, fmt.Sprintf("%s: %s", v[0], v[1]))
		}
		return fmt.Sprintf("%s: {%s}", v.Type().String(), strings.Join(sArr, ", "))
	default:
		return v.Type().String() + " value"
	}
}

func Display(name string, x interface{}) {
	fmt.Printf("Dsiplay %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), 0)
}

var limitDepth int = 3

func display(path string, v reflect.Value, depth int) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), depth)
		}
	case reflect.Struct:
		depth++
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			if depth <= limitDepth {
				display(fieldPath, v.Field(i), depth)
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key), depth)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), depth)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), depth)
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
