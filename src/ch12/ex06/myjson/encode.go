// Copyright (c) 2016 by akeboshi. All Rights Reserved.
package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func printWithKey(key, data string, buf *bytes.Buffer) {
	if key != "" {
		fmt.Fprintf(buf, "%s:%s", key, data)
	} else {
		fmt.Fprintf(buf, "%s", data)
	}
}

func encode(buf *bytes.Buffer, v reflect.Value, key string) error {
	switch v.Kind() {
	case reflect.Invalid:
		// not to do anything
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		if v.Int() != 0 {
			printWithKey(key, fmt.Sprintf("%d", v.Int()), buf)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		if v.Uint() != 0 {
			printWithKey(key, fmt.Sprintf("%d", v.Uint()), buf)
		}
	case reflect.String:
		printWithKey(key, fmt.Sprintf("%q", v.String()), buf)
	case reflect.Ptr:
		return encode(buf, v.Elem(), key)
	case reflect.Array, reflect.Slice:
		if len(key) > 2 {
			key = key[1 : len(key)-1]
		}
		if err := encode(buf, reflect.ValueOf(key), ""); err != nil {
			return err
		}
		buf.WriteByte(':')
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encode(buf, v.Index(i), ""); err != nil {
				return err
			}
		}
		buf.WriteByte(']')
	case reflect.Struct:
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 && v.NumField()-1 > i {
				buf.WriteByte(',')
			}
			key = fmt.Sprintf("%q", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), key); err != nil {
				return err
			}
		}
		buf.WriteByte('}')
	case reflect.Map:
		if err := encode(buf, reflect.ValueOf(key[1:len(key)-1]), ""); err != nil {
			return err
		}
		buf.WriteByte(':')
		buf.WriteByte('{')
		for i, kkey := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(',')
			}

			kkkey := fmt.Sprintf("%q", kkey.String())
			if err := encode(buf, v.MapIndex(kkey), kkkey); err != nil {
				return err
			}
		}
		buf.WriteByte('}')
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%e", v.Float())
	case reflect.Complex64, reflect.Complex128:
		val := v.Complex()
		fmt.Fprintf(buf, "#C(%f %f)", real(val), imag(val))
	case reflect.Bool:
		if v.Bool() {
			fmt.Fprintf(buf, "true")
		} else {
			fmt.Fprintf(buf, "false")
		}
	case reflect.Chan, reflect.Func:
		fmt.Fprintf(buf, "%s 0x%s", v.Type().String(), strconv.FormatUint(uint64(v.Pointer()), 16))
	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprintf(buf, "null")
		} else {
			fmt.Fprintf(buf, "\"%s\" ", v.Elem().Type())
			encode(buf, v.Elem(), key)
		}
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), ""); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
