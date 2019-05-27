package ex12_11

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func Pack(data interface{}) (url.URL, error) {
	// Build map of fields keyed by effective name
	fields := make(map[string]reflect.Value)
	v:= reflect.ValueOf(data).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	query := url.Values{}
	for key, val := range fields {
		if err := populate(query, key, val); err != nil {
			return url.URL{}, err
		}
	}

	return url.URL{RawQuery: query.Encode()}, nil
}

func populate(q url.Values, key string, v reflect.Value) error {
	switch v.Kind() {

	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			v := v.Field(i)
			key = fmt.Sprintf("%s[]", key)
			if err := populate(q, key, v); err != nil {
				return err
			}
		}

	case reflect.String:
		q.Add(key, v.String())

	case reflect.Bool:
		b := v.Bool()
		q.Add(key, strconv.FormatBool(b))

	case reflect.Int:
		i := v.Int()
		s := strconv.FormatInt(i, 10)
		q.Add(key, s)

	default:
		return fmt.Errorf("unsuported kind %s", v.Type())
	}

	return nil
}