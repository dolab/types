package types

import (
	"reflect"
	"time"
)

// IsEmpty returns true if v is of its type default value
func IsEmpty(v interface{}) bool {
	if IsZero(v) {
		return true
	}

	// spec for time.Time
	switch v.(type) {
	case time.Time:
		return v.(time.Time).IsZero()
	case *time.Time:
		return v.(*time.Time).IsZero()
	}

	var rv reflect.Value

	if tv, ok := v.(reflect.Value); ok {
		rv = tv
	} else {
		rv = reflect.ValueOf(v)
	}

	return !nonempty(rv)
}

func nonempty(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Field(i)
			if fv.Kind() == reflect.Interface && !fv.IsNil() {
				fv = fv.Elem()
			}
			if nonempty(fv) {
				return true
			}
		}
		return false
	case reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			if nonempty(rv.Index(i)) {
				return true
			}
		}
		return false
	case reflect.Slice, reflect.Map, reflect.Chan:
		return rv.Len() > 0
	case reflect.Ptr:
		return nonempty(rv.Elem())
	}
	return false
}
