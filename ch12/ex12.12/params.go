package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Check func(v interface{}) error

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}, checks map[string]Check) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		checkName := tag.Get("check")
		if checkName != "" {
			if check, ok := checks[checkName]; ok {
				if err := check(v.Field(i).Interface()); err != nil {
					return err
				}

			}
		}

		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}

		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}

	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}

	return nil
}