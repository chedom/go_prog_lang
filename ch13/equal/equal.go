package equal

import (
	"reflect"
	"unsafe"
)

type comparision struct {
	x, y unsafe.Pointer
	t reflect.Type
}

func equal(x, y reflect.Value, seen map[comparision]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}

	if x.Type() != y.Type() {
		return false
	}

	// cycle check
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(x.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparision{xptr, yptr, x.Type()}
		if seen[c] {
			return true //already seen
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	// ...numeric cases omitted for brevity

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}

		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	// ...struct and map cases omitted for brevity
	}
	panic("unreachable")
}

// Equal reports whether x and y are deeply equal.
func Equal(x, y interface{}) bool {
	seen := make(map[comparision]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}