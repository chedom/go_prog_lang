package display

import (
	"bytes"
	"fmt"
	"github.com/chedom/go_prog_lang/ch12/format"
	"reflect"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func formatMapKey(v reflect.Value) string {
	switch v.Kind() {

	case reflect.Array:
		buf := &bytes.Buffer{}
		buf.WriteByte('{')
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				buf.WriteByte(',')
			}
			buf.WriteString(format.FormatAtom(v.Index(i)))
		}
		buf.WriteByte('}')
		return buf.String()

	case reflect.Struct:
		buf := &bytes.Buffer{}
		buf.WriteByte('{')

		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				buf.WriteByte(',')
			}
			fmt.Fprintf(buf, "%s:%s", v.Type().Field(i).Name,
				format.FormatAtom(v.Field(i)))
		}

		buf.WriteByte('}')
		return buf.String()

	default:
		return format.FormatAtom(v)
	}
}

func display(path string, v reflect.Value, level int) {
	if level > 5 {
		fmt.Printf("%s = %s\n", path, format.FormatAtom(v))
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), level+1)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), level+1)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatMapKey(key)), v.MapIndex(key), level+1)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), level+1)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), level+1)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, format.FormatAtom(v))
	}
}
