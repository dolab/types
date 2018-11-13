package types

import "reflect"

// IsZero returns true if v is of its type default value
func IsZero(v interface{}) bool {
	if v == nil {
		return true
	}

	var rv reflect.Value

	if tv, ok := v.(reflect.Value); ok {
		rv = tv
	} else {
		rv = reflect.ValueOf(v)
	}
	if !rv.IsValid() {
		return true
	}

	if rv.Kind() == reflect.Interface && !rv.IsNil() {
		rv = rv.Elem()
	}

	return !nonzero(rv)
}

func nonzero(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Bool:
		return rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() != 0
	case reflect.Complex64, reflect.Complex128:
		return rv.Complex() != complex(0, 0)
	case reflect.String:
		return rv.String() != ""
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Field(i)
			if fv.Kind() == reflect.Interface && !fv.IsNil() {
				fv = fv.Elem()
			}
			if nonzero(fv) {
				return true
			}
		}
		return false
	case reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			if nonzero(rv.Index(i)) {
				return true
			}
		}
		return false
	case reflect.Interface, reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return !rv.IsNil()
	case reflect.UnsafePointer:
		return rv.Pointer() != 0
	}
	return true
}
