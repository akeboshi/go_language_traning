// Copyright (c) 2016 by akeboshi. All Rights Reserved.
package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func pTab(buf *bytes.Buffer, n int) {
	fmt.Fprintf(buf, "%*s", n, "\t")
}

func encode(buf *bytes.Buffer, v reflect.Value, depth int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Ptr:
		return encode(buf, v.Elem(), depth)
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		depth++
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte('\n')
				pTab(buf, depth)
			}
			if err := encode(buf, v.Index(i), depth); err != nil {
				return err
			}
		}
		depth--
		buf.WriteByte(')')
	case reflect.Struct:
		pTab(buf, depth)
		depth++
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s  ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), depth); err != nil {
				return err
			}
			buf.WriteByte(')')
			if i < v.NumField()-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')
	case reflect.Map:
		depth++
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte('\n')

				pTab(buf, depth)
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key, depth); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), depth); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
		depth--
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%e", v.Float())
	case reflect.Complex64, reflect.Complex128:
		val := v.Complex()
		fmt.Fprintf(buf, "#C(%f %f)", real(val), imag(val))
	case reflect.Bool:
		if v.Bool() {
			fmt.Fprintf(buf, "t")
		} else {
			fmt.Fprintf(buf, "nil")
		}
	case reflect.Chan, reflect.Func:
		fmt.Fprintf(buf, "%s 0x%s", v.Type().String(), strconv.FormatUint(uint64(v.Pointer()), 16))
	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprintf(buf, "nil")
		} else {
			fmt.Fprintf(buf, "\"%s\" ", v.Elem().Type())
			encode(buf, v.Elem(), depth)
		}
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
