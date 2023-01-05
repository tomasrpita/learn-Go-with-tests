package reflection

import "reflect"

// Walk takes a struct x and calls fn for all strings fields found inside.
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// walkValue is a closure that calls walk with the value of the field.
	walkVaue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkVaue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkVaue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkVaue(val.MapIndex(key))
		}

	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkVaue(v)
		}
	case reflect.Func:
		valResult := val.Call(nil)
		for _, res := range valResult {
			walkVaue(res)
		}
	}

}

// getValue returns the value of x. If x is a pointer, it returns the value.
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
