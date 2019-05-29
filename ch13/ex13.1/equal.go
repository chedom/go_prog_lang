package ex13_1

import "reflect"

// Consider numbers equalish if their difference is less than one part in
// <multiplier>.
const multiplier = 1000000000

func Equal(x, y interface{}) bool {
	return equal(reflect.ValueOf(x), reflect.ValueOf(y))
}

func equal(x, y reflect.Value) bool {
	if x.Type() != y.Type() {
		return false
	}

	switch x.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return equalNum(float64(x.Uint()), float64(y.Uint()))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return equalNum(float64(x.Int()), float64(y.Int()))

	case reflect.Float32, reflect.Float64:
		return equalNum(float64(x.Float()), float64(y.Float()))

	case reflect.Complex64, reflect.Complex128:
		realEqual := equalNum(float64(real(x.Complex())), float64(real(y.Complex())))
		imagEqual := equalNum(float64(imag(x.Complex())), float64(imag(y.Complex())))
		return realEqual && imagEqual

	default:
		return false
	}

}


func equalNum(x, y float64) bool {
	if x == y {
		return true
	}

	var diff float64
	if x > y {
		diff = x - y
	} else {
		diff = y - x
	}

	d := diff * multiplier
	return (d < x) && (d < y)

}